package pub

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
	"net/http"
	"net/url"
	"strings"
)

var (
	// ErrObjectRequired indicates the activity needs its object property
	// set. Can be returned by DelegateActor's PostInbox or PostOutbox so a
	// Bad Request response is set.
	ErrObjectRequired = errors.New("object property required on the provided activity")
	// ErrTargetRequired indicates the activity needs its target property
	// set. Can be returned by DelegateActor's PostInbox or PostOutbox so a
	// Bad Request response is set.
	ErrTargetRequired = errors.New("target property required on the provided activity")
)

// activityStreamsMediaTypes contains all of the accepted ActivityStreams media
// types. Generated at init time.
var activityStreamsMediaTypes []string

func init() {
	activityStreamsMediaTypes = []string{
		"application/activity+json",
	}
	jsonLdType := "application/ld+json"
	for _, semi := range []string{";", " ;", " ; ", "; "} {
		for _, profile := range []string{
			"profile=https://www.w3.org/ns/activitystreams",
			"profile=\"https://www.w3.org/ns/activitystreams\"",
		} {
			activityStreamsMediaTypes = append(
				activityStreamsMediaTypes,
				fmt.Sprintf("%s%s%s", jsonLdType, semi, profile))
		}
	}
}

// headerIsActivityPubMediaType returns true if the header string contains one
// of the accepted ActivityStreams media types.
//
// Note we don't try to build a comprehensive parser and instead accept a
// tolerable amount of whitespace since the HTTP specification is ambiguous
// about the format and significance of whitespace.
func headerIsActivityPubMediaType(header string) bool {
	for _, mediaType := range activityStreamsMediaTypes {
		if strings.Contains(header, mediaType) {
			return true
		}
	}
	return false
}

const (
	// The Content-Type header.
	contentTypeHeader = "Content-Type"
	// The Accept header.
	acceptHeader = "Accept"
)

// isActivityPubPost returns true if the request is a POST request that has the
// ActivityStreams content type header
func isActivityPubPost(r *http.Request) bool {
	return r.Method == "POST" && headerIsActivityPubMediaType(r.Header.Get(contentTypeHeader))
}

// isActivityPubGet returns true if the request is a GET request that has the
// ActivityStreams content type header
func isActivityPubGet(r *http.Request) bool {
	return r.Method == "GET" && headerIsActivityPubMediaType(r.Header.Get(acceptHeader))
}

// dedupeOrderedItems deduplicates the 'orderedItems' within an ordered
// collection type. Deduplication happens by the 'id' property.
func dedupeOrderedItems(oc orderedItemser) error {
	oi := oc.GetActivityStreamsOrderedItems()
	if oi == nil {
		return nil
	}
	seen := make(map[string]bool, oi.Len())
	for i := 0; i < oi.Len(); {
		var id *url.URL

		iter := oi.At(i)
		asType := iter.GetType()
		if asType != nil {
			var err error
			id, err = GetId(asType)
			if err != nil {
				return err
			}
		} else if iter.IsIRI() {
			id = iter.GetIRI()
		} else {
			return fmt.Errorf("element %d in OrderedCollection does not have an ID nor is an IRI", i)
		}
		if seen[id.String()] {
			oi.Remove(i)
		} else {
			seen[id.String()] = true
			i++
		}
	}
	return nil
}

const (
	// jsonLDContext is the key for the JSON-LD specification's context
	// value. It contains the definitions of the types contained within the
	// rest of the payload. Important for linked-data representations, but
	// only applicable to go-fed at code-generation time.
	jsonLDContext = "@context"
)

// addJSONLDContext adds the
func serialize(a vocab.Type) (m map[string]interface{}, e error) {
	m, e = a.Serialize()
	if e != nil {
		return
	}
	v := a.JSONLDContext()
	// Transform the map of vocabulary-to-aliases into a context payload,
	// but do so in a way that at least keeps it readable for other humans.
	var contextValue interface{}
	if len(v) == 1 {
		for vocab, alias := range v {
			if len(alias) == 0 {
				contextValue = vocab
			} else {
				contextValue = map[string]string{
					alias: vocab,
				}
			}
		}
	} else {
		var arr []interface{}
		aliases := make(map[string]string)
		for vocab, alias := range v {
			if len(alias) == 0 {
				arr = append(arr, vocab)
			} else {
				aliases[alias] = vocab
			}
		}
		contextValue = append(arr, aliases)
	}
	m[jsonLDContext] = contextValue
	return
}

const (
	// Contains the ActivityStreams Content-Type value.
	contentTypeHeaderValue = "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""
	// The Date header.
	dateHeader = "Date"
	// The Digest header.
	digestHeader = "Digest"
	// The delimiter used in the Digest header.
	digestDelimiter = "="
	// SHA-256 string for the Digest header.
	sha256Digest = "SHA-256"
)

// addResponseHeaders sets headers needed in the HTTP response, such but not
// limited to the Content-Type, Date, and Digest headers.
func addResponseHeaders(h http.Header, c Clock, responseContent []byte) {
	h.Set(contentTypeHeader, contentTypeHeaderValue)
	// RFC 7231 §7.1.1.2
	h.Set(dateHeader, c.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	// RFC 3230 and RFC 5843
	var b bytes.Buffer
	b.WriteString(sha256Digest)
	b.WriteString(digestDelimiter)
	hashed := sha256.Sum256(responseContent)
	b.WriteString(base64.StdEncoding.EncodeToString(hashed[:]))
	h.Set(digestHeader, b.String())
}

// IdProperty is a property that can readily have its id obtained
type IdProperty interface {
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() vocab.Type
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
}

// ToId returns an IdProperty's id.
func ToId(i IdProperty) (*url.URL, error) {
	if i.GetType() != nil {
		return GetId(i.GetType())
	} else if i.IsIRI() {
		return i.GetIRI(), nil
	}
	return nil, fmt.Errorf("cannot determine id of activitystreams property")
}

// GetId will attempt to find the 'id' property or, if it happens to be a
// Link or derived from Link type, the 'href' property instead.
//
// Returns an error if the id is not set and either the 'href' property is not
// valid on this type, or it is also not set.
func GetId(t vocab.Type) (*url.URL, error) {
	if id := t.GetActivityStreamsId(); id != nil {
		return id.Get(), nil
	} else if h, ok := t.(hrefer); ok {
		if href := h.GetActivityStreamsHref(); href != nil {
			return href.Get(), nil
		}
	}
	return nil, fmt.Errorf("cannot determine id of activitystreams value")
}

// getInboxForwardingValues obtains the 'inReplyTo', 'object', 'target', and
// 'tag' values on an ActivityStreams value.
func getInboxForwardingValues(o vocab.Type) (t []vocab.Type, iri []*url.URL) {
	// 'inReplyTo'
	if i, ok := o.(inReplyToer); ok {
		irt := i.GetActivityStreamsInReplyTo()
		for iter := irt.Begin(); iter != irt.End(); iter = iter.Next() {
			if tv := iter.GetType(); tv != nil {
				t = append(t, tv)
			} else {
				iri = append(iri, iter.GetIRI())
			}
		}
	}
	// 'tag'
	if i, ok := o.(tagger); ok {
		tag := i.GetActivityStreamsTag()
		for iter := tag.Begin(); iter != tag.End(); iter = iter.Next() {
			if tv := iter.GetType(); tv != nil {
				t = append(t, tv)
			} else {
				iri = append(iri, iter.GetIRI())
			}
		}
	}
	// 'object'
	if i, ok := o.(objecter); ok {
		obj := i.GetActivityStreamsObject()
		for iter := obj.Begin(); iter != obj.End(); iter = iter.Next() {
			if tv := iter.GetType(); tv != nil {
				t = append(t, tv)
			} else {
				iri = append(iri, iter.GetIRI())
			}
		}
	}
	// 'target'
	if i, ok := o.(targeter); ok {
		tar := i.GetActivityStreamsTarget()
		for iter := tar.Begin(); iter != tar.End(); iter = iter.Next() {
			if tv := iter.GetType(); tv != nil {
				t = append(t, tv)
			} else {
				iri = append(iri, iter.GetIRI())
			}
		}
	}
	return
}

// wrapInCreate will automatically wrap the provided object in a Create
// activity. This will copy over the 'to', 'bto', 'cc', 'bcc', and 'audience'
// properties. It will also copy over the published time if present.
func wrapInCreate(ctx context.Context, o vocab.Type, actor *url.URL) (c vocab.ActivityStreamsCreate, err error) {
	c = streams.NewActivityStreamsCreate()
	// Object property
	oProp := streams.NewActivityStreamsObjectProperty()
	addToCreate(ctx, c, o)
	c.SetActivityStreamsObject(oProp)
	// Actor Property
	actorProp := streams.NewActivityStreamsActorProperty()
	actorProp.AppendIRI(actor)
	c.SetActivityStreamsActor(actorProp)
	// Published Property
	if v, ok := o.(publisheder); ok {
		c.SetActivityStreamsPublished(v.GetActivityStreamsPublished())
	}
	// Copying over properties.
	if v, ok := o.(toer); ok {
		activityTo := streams.NewActivityStreamsToProperty()
		to := v.GetActivityStreamsTo()
		for iter := to.Begin(); iter != to.End(); iter = iter.Next() {
			var id *url.URL
			id, err = ToId(iter)
			if err != nil {
				return
			}
			activityTo.AppendIRI(id)
		}
		c.SetActivityStreamsTo(activityTo)
	}
	if v, ok := o.(btoer); ok {
		activityBto := streams.NewActivityStreamsBtoProperty()
		bto := v.GetActivityStreamsBto()
		for iter := bto.Begin(); iter != bto.End(); iter = iter.Next() {
			var id *url.URL
			id, err = ToId(iter)
			if err != nil {
				return
			}
			activityBto.AppendIRI(id)
		}
		c.SetActivityStreamsBto(activityBto)
	}
	if v, ok := o.(ccer); ok {
		activityCc := streams.NewActivityStreamsCcProperty()
		cc := v.GetActivityStreamsCc()
		for iter := cc.Begin(); iter != cc.End(); iter = iter.Next() {
			var id *url.URL
			id, err = ToId(iter)
			if err != nil {
				return
			}
			activityCc.AppendIRI(id)
		}
		c.SetActivityStreamsCc(activityCc)
	}
	if v, ok := o.(bccer); ok {
		activityBcc := streams.NewActivityStreamsBccProperty()
		bcc := v.GetActivityStreamsBcc()
		for iter := bcc.Begin(); iter != bcc.End(); iter = iter.Next() {
			var id *url.URL
			id, err = ToId(iter)
			if err != nil {
				return
			}
			activityBcc.AppendIRI(id)
		}
		c.SetActivityStreamsBcc(activityBcc)
	}
	if v, ok := o.(audiencer); ok {
		activityAudience := streams.NewActivityStreamsAudienceProperty()
		aud := v.GetActivityStreamsAudience()
		for iter := aud.Begin(); iter != aud.End(); iter = iter.Next() {
			var id *url.URL
			id, err = ToId(iter)
			if err != nil {
				return
			}
			activityAudience.AppendIRI(id)
		}
		c.SetActivityStreamsAudience(activityAudience)
	}
	return
}

// filterURLs removes urls whose strings match the provided filter
func filterURLs(u []*url.URL, fn func(s string) bool) []*url.URL {
	i := 0
	for i < len(u) {
		if fn(u[i].String()) {
			u = append(u[:i], u[i+1:]...)
		} else {
			i++
		}
	}
	return u
}

const (
	// PublicActivityPubIRI is the IRI that indicates an Activity is meant
	// to be visible for general public consumption.
	PublicActivityPubIRI = "https://www.w3.org/ns/activitystreams#Public"
	publicJsonLD         = "Public"
	publicJsonLDAS       = "as:Public"
)

// IsPublic determines if an IRI string is the Public collection as defined in
// the spec, including JSON-LD compliant collections.
func IsPublic(s string) bool {
	return s == PublicActivityPubIRI || s == publicJsonLD || s == publicJsonLDAS
}

// getInboxes extracts the 'inbox' IRIs from actor types.
func getInboxes(t []vocab.Type) (u []*url.URL, err error) {
	for _, elem := range t {
		var iri *url.URL
		iri, err = getInbox(elem)
		if err != nil {
			return
		}
		u = append(u, iri)
	}
	return
}

// getInbox extracts the 'inbox' IRI from an actor type.
func getInbox(t vocab.Type) (u *url.URL, err error) {
	ib, ok := t.(inboxer)
	if !ok {
		err = fmt.Errorf("actor type %T has no inbox", t)
		return
	}
	inbox := ib.GetActivityStreamsInbox()
	return ToId(inbox)
}

// dedupeIRIs will deduplicate final inbox IRIs. The ignore list is applied to
// the final list.
func dedupeIRIs(recipients, ignored []*url.URL) (out []*url.URL) {
	ignoredMap := make(map[string]bool, len(ignored))
	for _, elem := range ignored {
		ignoredMap[elem.String()] = true
	}
	outMap := make(map[string]bool, len(recipients))
	for _, k := range recipients {
		kStr := k.String()
		if !ignoredMap[kStr] && !outMap[kStr] {
			out = append(out, k)
			outMap[kStr] = true
		}
	}
	return
}

// stripHiddenRecipients removes "bto" and "bcc" from the activity.
//
// Note that this requirement of the specification is under "Section 6: Client
// to Server Interactions", the Social API, and not the Federative API.
func stripHiddenRecipients(activity Activity) {
	bto := activity.GetActivityStreamsBto()
	if bto != nil {
		for i := bto.Len() - 1; i >= 0; i-- {
			bto.Remove(i)
		}
	}
	bcc := activity.GetActivityStreamsBcc()
	if bcc != nil {
		for i := bcc.Len() - 1; i >= 0; i-- {
			bcc.Remove(i)
		}
	}
}
