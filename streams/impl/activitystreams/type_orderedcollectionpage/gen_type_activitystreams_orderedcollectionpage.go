package typeorderedcollectionpage

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"strings"
)

// Used to represent ordered subsets of items from an OrderedCollection. Refer to
// the Activity Streams 2.0 Core for a complete description of the
// OrderedCollectionPage object.
//
// Example 8 (https://www.w3.org/TR/activitystreams-vocabulary/#ex6c-jsonld):
//   {
//     "id": "http://example.org/foo?page=1",
//     "orderedItems": [
//       {
//         "name": "A Simple Note",
//         "type": "Note"
//       },
//       {
//         "name": "Another Simple Note",
//         "type": "Note"
//       }
//     ],
//     "partOf": "http://example.org/foo",
//     "summary": "Page 1 of Sally's notes",
//     "type": "OrderedCollectionPage"
//   }
type ActivityStreamsOrderedCollectionPage struct {
	ActivityStreamsAltitude     vocab.ActivityStreamsAltitudeProperty
	ActivityStreamsAttachment   vocab.ActivityStreamsAttachmentProperty
	ActivityStreamsAttributedTo vocab.ActivityStreamsAttributedToProperty
	ActivityStreamsAudience     vocab.ActivityStreamsAudienceProperty
	ActivityStreamsBcc          vocab.ActivityStreamsBccProperty
	ActivityStreamsBto          vocab.ActivityStreamsBtoProperty
	ActivityStreamsCc           vocab.ActivityStreamsCcProperty
	ActivityStreamsContent      vocab.ActivityStreamsContentProperty
	ActivityStreamsContext      vocab.ActivityStreamsContextProperty
	ActivityStreamsCurrent      vocab.ActivityStreamsCurrentProperty
	ActivityStreamsDuration     vocab.ActivityStreamsDurationProperty
	ActivityStreamsEndTime      vocab.ActivityStreamsEndTimeProperty
	ActivityStreamsFirst        vocab.ActivityStreamsFirstProperty
	ActivityStreamsGenerator    vocab.ActivityStreamsGeneratorProperty
	ActivityStreamsIcon         vocab.ActivityStreamsIconProperty
	ActivityStreamsId           vocab.ActivityStreamsIdProperty
	ActivityStreamsImage        vocab.ActivityStreamsImageProperty
	ActivityStreamsInReplyTo    vocab.ActivityStreamsInReplyToProperty
	ActivityStreamsLast         vocab.ActivityStreamsLastProperty
	ActivityStreamsLikes        vocab.ActivityStreamsLikesProperty
	ActivityStreamsLocation     vocab.ActivityStreamsLocationProperty
	ActivityStreamsMediaType    vocab.ActivityStreamsMediaTypeProperty
	ActivityStreamsName         vocab.ActivityStreamsNameProperty
	ActivityStreamsNext         vocab.ActivityStreamsNextProperty
	ActivityStreamsObject       vocab.ActivityStreamsObjectProperty
	ActivityStreamsOrderedItems vocab.ActivityStreamsOrderedItemsProperty
	ActivityStreamsPartOf       vocab.ActivityStreamsPartOfProperty
	ActivityStreamsPrev         vocab.ActivityStreamsPrevProperty
	ActivityStreamsPreview      vocab.ActivityStreamsPreviewProperty
	ActivityStreamsPublished    vocab.ActivityStreamsPublishedProperty
	ActivityStreamsReplies      vocab.ActivityStreamsRepliesProperty
	ActivityStreamsShares       vocab.ActivityStreamsSharesProperty
	ActivityStreamsStartIndex   vocab.ActivityStreamsStartIndexProperty
	ActivityStreamsStartTime    vocab.ActivityStreamsStartTimeProperty
	ActivityStreamsSummary      vocab.ActivityStreamsSummaryProperty
	ActivityStreamsTag          vocab.ActivityStreamsTagProperty
	ActivityStreamsTo           vocab.ActivityStreamsToProperty
	ActivityStreamsTotalItems   vocab.ActivityStreamsTotalItemsProperty
	ActivityStreamsType         vocab.ActivityStreamsTypeProperty
	ActivityStreamsUpdated      vocab.ActivityStreamsUpdatedProperty
	ActivityStreamsUrl          vocab.ActivityStreamsUrlProperty
	alias                       string
	unknown                     map[string]interface{}
}

// ActivityStreamsOrderedCollectionPageExtends returns true if the
// OrderedCollectionPage type extends from the other type.
func ActivityStreamsOrderedCollectionPageExtends(other vocab.Type) bool {
	extensions := []string{"Collection", "CollectionPage", "Object", "OrderedCollection"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// DeserializeOrderedCollectionPage creates a OrderedCollectionPage from a map
// representation that has been unmarshalled from a text or binary format.
func DeserializeOrderedCollectionPage(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsOrderedCollectionPage, error) {
	alias := ""
	aliasPrefix := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
		aliasPrefix = a + ":"
	}
	this := &ActivityStreamsOrderedCollectionPage{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	if typeValue, ok := m["type"]; !ok {
		return nil, fmt.Errorf("no \"type\" property in map")
	} else if typeString, ok := typeValue.(string); ok {
		typeName := strings.TrimPrefix(typeString, aliasPrefix)
		if typeName != "OrderedCollectionPage" {
			return nil, fmt.Errorf("\"type\" property is not of %q type: %s", "OrderedCollectionPage", typeName)
		}
		// Fall through, success in finding a proper Type
	} else if arrType, ok := typeValue.([]interface{}); ok {
		found := false
		for _, elemVal := range arrType {
			if typeString, ok := elemVal.(string); ok && strings.TrimPrefix(typeString, aliasPrefix) == "OrderedCollectionPage" {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("could not find a \"type\" property of value %q", "OrderedCollectionPage")
		}
		// Fall through, success in finding a proper Type
	} else {
		return nil, fmt.Errorf("\"type\" property is unrecognized type: %T", typeValue)
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeAltitudePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAltitude = p
	}
	if p, err := mgr.DeserializeAttachmentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAttachment = p
	}
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAttributedTo = p
	}
	if p, err := mgr.DeserializeAudiencePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAudience = p
	}
	if p, err := mgr.DeserializeBccPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsBcc = p
	}
	if p, err := mgr.DeserializeBtoPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsBto = p
	}
	if p, err := mgr.DeserializeCcPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsCc = p
	}
	if p, err := mgr.DeserializeContentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsContent = p
	}
	if p, err := mgr.DeserializeContextPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsContext = p
	}
	if p, err := mgr.DeserializeCurrentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsCurrent = p
	}
	if p, err := mgr.DeserializeDurationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsDuration = p
	}
	if p, err := mgr.DeserializeEndTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsEndTime = p
	}
	if p, err := mgr.DeserializeFirstPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsFirst = p
	}
	if p, err := mgr.DeserializeGeneratorPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsGenerator = p
	}
	if p, err := mgr.DeserializeIconPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsIcon = p
	}
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsId = p
	}
	if p, err := mgr.DeserializeImagePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsImage = p
	}
	if p, err := mgr.DeserializeInReplyToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsInReplyTo = p
	}
	if p, err := mgr.DeserializeLastPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsLast = p
	}
	if p, err := mgr.DeserializeLikesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsLikes = p
	}
	if p, err := mgr.DeserializeLocationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsLocation = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsMediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsName = p
	}
	if p, err := mgr.DeserializeNextPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsNext = p
	}
	if p, err := mgr.DeserializeObjectPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsObject = p
	}
	if p, err := mgr.DeserializeOrderedItemsPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsOrderedItems = p
	}
	if p, err := mgr.DeserializePartOfPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPartOf = p
	}
	if p, err := mgr.DeserializePrevPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPrev = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPreview = p
	}
	if p, err := mgr.DeserializePublishedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPublished = p
	}
	if p, err := mgr.DeserializeRepliesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsReplies = p
	}
	if p, err := mgr.DeserializeSharesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsShares = p
	}
	if p, err := mgr.DeserializeStartIndexPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsStartIndex = p
	}
	if p, err := mgr.DeserializeStartTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsStartTime = p
	}
	if p, err := mgr.DeserializeSummaryPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsSummary = p
	}
	if p, err := mgr.DeserializeTagPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsTag = p
	}
	if p, err := mgr.DeserializeToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsTo = p
	}
	if p, err := mgr.DeserializeTotalItemsPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsTotalItems = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsType = p
	}
	if p, err := mgr.DeserializeUpdatedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsUpdated = p
	}
	if p, err := mgr.DeserializeUrlPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsUrl = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "altitude" {
			continue
		} else if k == "attachment" {
			continue
		} else if k == "attributedTo" {
			continue
		} else if k == "audience" {
			continue
		} else if k == "bcc" {
			continue
		} else if k == "bto" {
			continue
		} else if k == "cc" {
			continue
		} else if k == "content" {
			continue
		} else if k == "contentMap" {
			continue
		} else if k == "context" {
			continue
		} else if k == "current" {
			continue
		} else if k == "duration" {
			continue
		} else if k == "endTime" {
			continue
		} else if k == "first" {
			continue
		} else if k == "generator" {
			continue
		} else if k == "icon" {
			continue
		} else if k == "id" {
			continue
		} else if k == "image" {
			continue
		} else if k == "inReplyTo" {
			continue
		} else if k == "last" {
			continue
		} else if k == "likes" {
			continue
		} else if k == "location" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "nameMap" {
			continue
		} else if k == "next" {
			continue
		} else if k == "object" {
			continue
		} else if k == "orderedItems" {
			continue
		} else if k == "partOf" {
			continue
		} else if k == "prev" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "published" {
			continue
		} else if k == "replies" {
			continue
		} else if k == "shares" {
			continue
		} else if k == "startIndex" {
			continue
		} else if k == "startTime" {
			continue
		} else if k == "summary" {
			continue
		} else if k == "summaryMap" {
			continue
		} else if k == "tag" {
			continue
		} else if k == "to" {
			continue
		} else if k == "totalItems" {
			continue
		} else if k == "type" {
			continue
		} else if k == "updated" {
			continue
		} else if k == "url" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// NewActivityStreamsOrderedCollectionPage creates a new OrderedCollectionPage type
func NewActivityStreamsOrderedCollectionPage() *ActivityStreamsOrderedCollectionPage {
	typeProp := typePropertyConstructor()
	typeProp.AppendXMLSchemaString("OrderedCollectionPage")
	return &ActivityStreamsOrderedCollectionPage{
		ActivityStreamsType: typeProp,
		alias:               "",
		unknown:             make(map[string]interface{}, 0),
	}
}

// OrderedCollectionPageIsDisjointWith returns true if the other provided type is
// disjoint with the OrderedCollectionPage type.
func OrderedCollectionPageIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Link", "Mention"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// OrderedCollectionPageIsExtendedBy returns true if the other provided type
// extends from the OrderedCollectionPage type.
func OrderedCollectionPageIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// GetActivityStreamsAltitude returns the "altitude" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsAltitude() vocab.ActivityStreamsAltitudeProperty {
	return this.ActivityStreamsAltitude
}

// GetActivityStreamsAttachment returns the "attachment" property if it exists,
// and nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsAttachment() vocab.ActivityStreamsAttachmentProperty {
	return this.ActivityStreamsAttachment
}

// GetActivityStreamsAttributedTo returns the "attributedTo" property if it
// exists, and nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsAttributedTo() vocab.ActivityStreamsAttributedToProperty {
	return this.ActivityStreamsAttributedTo
}

// GetActivityStreamsAudience returns the "audience" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsAudience() vocab.ActivityStreamsAudienceProperty {
	return this.ActivityStreamsAudience
}

// GetActivityStreamsBcc returns the "bcc" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsBcc() vocab.ActivityStreamsBccProperty {
	return this.ActivityStreamsBcc
}

// GetActivityStreamsBto returns the "bto" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsBto() vocab.ActivityStreamsBtoProperty {
	return this.ActivityStreamsBto
}

// GetActivityStreamsCc returns the "cc" property if it exists, and nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsCc() vocab.ActivityStreamsCcProperty {
	return this.ActivityStreamsCc
}

// GetActivityStreamsContent returns the "content" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsContent() vocab.ActivityStreamsContentProperty {
	return this.ActivityStreamsContent
}

// GetActivityStreamsContext returns the "context" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsContext() vocab.ActivityStreamsContextProperty {
	return this.ActivityStreamsContext
}

// GetActivityStreamsCurrent returns the "current" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsCurrent() vocab.ActivityStreamsCurrentProperty {
	return this.ActivityStreamsCurrent
}

// GetActivityStreamsDuration returns the "duration" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsDuration() vocab.ActivityStreamsDurationProperty {
	return this.ActivityStreamsDuration
}

// GetActivityStreamsEndTime returns the "endTime" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsEndTime() vocab.ActivityStreamsEndTimeProperty {
	return this.ActivityStreamsEndTime
}

// GetActivityStreamsFirst returns the "first" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsFirst() vocab.ActivityStreamsFirstProperty {
	return this.ActivityStreamsFirst
}

// GetActivityStreamsGenerator returns the "generator" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsGenerator() vocab.ActivityStreamsGeneratorProperty {
	return this.ActivityStreamsGenerator
}

// GetActivityStreamsIcon returns the "icon" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsIcon() vocab.ActivityStreamsIconProperty {
	return this.ActivityStreamsIcon
}

// GetActivityStreamsId returns the "id" property if it exists, and nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsId() vocab.ActivityStreamsIdProperty {
	return this.ActivityStreamsId
}

// GetActivityStreamsImage returns the "image" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsImage() vocab.ActivityStreamsImageProperty {
	return this.ActivityStreamsImage
}

// GetActivityStreamsInReplyTo returns the "inReplyTo" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsInReplyTo() vocab.ActivityStreamsInReplyToProperty {
	return this.ActivityStreamsInReplyTo
}

// GetActivityStreamsLast returns the "last" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsLast() vocab.ActivityStreamsLastProperty {
	return this.ActivityStreamsLast
}

// GetActivityStreamsLikes returns the "likes" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsLikes() vocab.ActivityStreamsLikesProperty {
	return this.ActivityStreamsLikes
}

// GetActivityStreamsLocation returns the "location" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsLocation() vocab.ActivityStreamsLocationProperty {
	return this.ActivityStreamsLocation
}

// GetActivityStreamsMediaType returns the "mediaType" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsMediaType() vocab.ActivityStreamsMediaTypeProperty {
	return this.ActivityStreamsMediaType
}

// GetActivityStreamsName returns the "name" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsName() vocab.ActivityStreamsNameProperty {
	return this.ActivityStreamsName
}

// GetActivityStreamsNext returns the "next" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsNext() vocab.ActivityStreamsNextProperty {
	return this.ActivityStreamsNext
}

// GetActivityStreamsObject returns the "object" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsObject() vocab.ActivityStreamsObjectProperty {
	return this.ActivityStreamsObject
}

// GetActivityStreamsOrderedItems returns the "orderedItems" property if it
// exists, and nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsOrderedItems() vocab.ActivityStreamsOrderedItemsProperty {
	return this.ActivityStreamsOrderedItems
}

// GetActivityStreamsPartOf returns the "partOf" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsPartOf() vocab.ActivityStreamsPartOfProperty {
	return this.ActivityStreamsPartOf
}

// GetActivityStreamsPrev returns the "prev" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsPrev() vocab.ActivityStreamsPrevProperty {
	return this.ActivityStreamsPrev
}

// GetActivityStreamsPreview returns the "preview" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsPreview() vocab.ActivityStreamsPreviewProperty {
	return this.ActivityStreamsPreview
}

// GetActivityStreamsPublished returns the "published" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsPublished() vocab.ActivityStreamsPublishedProperty {
	return this.ActivityStreamsPublished
}

// GetActivityStreamsReplies returns the "replies" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsReplies() vocab.ActivityStreamsRepliesProperty {
	return this.ActivityStreamsReplies
}

// GetActivityStreamsShares returns the "shares" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsShares() vocab.ActivityStreamsSharesProperty {
	return this.ActivityStreamsShares
}

// GetActivityStreamsStartIndex returns the "startIndex" property if it exists,
// and nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsStartIndex() vocab.ActivityStreamsStartIndexProperty {
	return this.ActivityStreamsStartIndex
}

// GetActivityStreamsStartTime returns the "startTime" property if it exists, and
// nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsStartTime() vocab.ActivityStreamsStartTimeProperty {
	return this.ActivityStreamsStartTime
}

// GetActivityStreamsSummary returns the "summary" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsSummary() vocab.ActivityStreamsSummaryProperty {
	return this.ActivityStreamsSummary
}

// GetActivityStreamsTag returns the "tag" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsTag() vocab.ActivityStreamsTagProperty {
	return this.ActivityStreamsTag
}

// GetActivityStreamsTo returns the "to" property if it exists, and nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsTo() vocab.ActivityStreamsToProperty {
	return this.ActivityStreamsTo
}

// GetActivityStreamsTotalItems returns the "totalItems" property if it exists,
// and nil otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsTotalItems() vocab.ActivityStreamsTotalItemsProperty {
	return this.ActivityStreamsTotalItems
}

// GetActivityStreamsType returns the "type" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsType() vocab.ActivityStreamsTypeProperty {
	return this.ActivityStreamsType
}

// GetActivityStreamsUpdated returns the "updated" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsUpdated() vocab.ActivityStreamsUpdatedProperty {
	return this.ActivityStreamsUpdated
}

// GetActivityStreamsUrl returns the "url" property if it exists, and nil
// otherwise.
func (this ActivityStreamsOrderedCollectionPage) GetActivityStreamsUrl() vocab.ActivityStreamsUrlProperty {
	return this.ActivityStreamsUrl
}

// GetName returns the name of this type.
func (this ActivityStreamsOrderedCollectionPage) GetName() string {
	return "OrderedCollectionPage"
}

// GetUnknownProperties returns the unknown properties for the
// OrderedCollectionPage type. Note that this should not be used by app
// developers. It is only used to help determine which implementation is
// LessThan the other. Developers who are creating a different implementation
// of this type's interface can use this method in their LessThan
// implementation, but routine ActivityPub applications should not use this to
// bypass the code generation tool.
func (this ActivityStreamsOrderedCollectionPage) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// IsExtending returns true if the OrderedCollectionPage type extends from the
// other type.
func (this ActivityStreamsOrderedCollectionPage) IsExtending(other vocab.Type) bool {
	return ActivityStreamsOrderedCollectionPageExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this ActivityStreamsOrderedCollectionPage) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	m = this.helperJSONLDContext(this.ActivityStreamsAltitude, m)
	m = this.helperJSONLDContext(this.ActivityStreamsAttachment, m)
	m = this.helperJSONLDContext(this.ActivityStreamsAttributedTo, m)
	m = this.helperJSONLDContext(this.ActivityStreamsAudience, m)
	m = this.helperJSONLDContext(this.ActivityStreamsBcc, m)
	m = this.helperJSONLDContext(this.ActivityStreamsBto, m)
	m = this.helperJSONLDContext(this.ActivityStreamsCc, m)
	m = this.helperJSONLDContext(this.ActivityStreamsContent, m)
	m = this.helperJSONLDContext(this.ActivityStreamsContext, m)
	m = this.helperJSONLDContext(this.ActivityStreamsCurrent, m)
	m = this.helperJSONLDContext(this.ActivityStreamsDuration, m)
	m = this.helperJSONLDContext(this.ActivityStreamsEndTime, m)
	m = this.helperJSONLDContext(this.ActivityStreamsFirst, m)
	m = this.helperJSONLDContext(this.ActivityStreamsGenerator, m)
	m = this.helperJSONLDContext(this.ActivityStreamsIcon, m)
	m = this.helperJSONLDContext(this.ActivityStreamsId, m)
	m = this.helperJSONLDContext(this.ActivityStreamsImage, m)
	m = this.helperJSONLDContext(this.ActivityStreamsInReplyTo, m)
	m = this.helperJSONLDContext(this.ActivityStreamsLast, m)
	m = this.helperJSONLDContext(this.ActivityStreamsLikes, m)
	m = this.helperJSONLDContext(this.ActivityStreamsLocation, m)
	m = this.helperJSONLDContext(this.ActivityStreamsMediaType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsName, m)
	m = this.helperJSONLDContext(this.ActivityStreamsNext, m)
	m = this.helperJSONLDContext(this.ActivityStreamsObject, m)
	m = this.helperJSONLDContext(this.ActivityStreamsOrderedItems, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPartOf, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPrev, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPreview, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPublished, m)
	m = this.helperJSONLDContext(this.ActivityStreamsReplies, m)
	m = this.helperJSONLDContext(this.ActivityStreamsShares, m)
	m = this.helperJSONLDContext(this.ActivityStreamsStartIndex, m)
	m = this.helperJSONLDContext(this.ActivityStreamsStartTime, m)
	m = this.helperJSONLDContext(this.ActivityStreamsSummary, m)
	m = this.helperJSONLDContext(this.ActivityStreamsTag, m)
	m = this.helperJSONLDContext(this.ActivityStreamsTo, m)
	m = this.helperJSONLDContext(this.ActivityStreamsTotalItems, m)
	m = this.helperJSONLDContext(this.ActivityStreamsType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsUpdated, m)
	m = this.helperJSONLDContext(this.ActivityStreamsUrl, m)

	return m
}

// LessThan computes if this OrderedCollectionPage is lesser, with an arbitrary
// but stable determination.
func (this ActivityStreamsOrderedCollectionPage) LessThan(o vocab.ActivityStreamsOrderedCollectionPage) bool {
	// Begin: Compare known properties
	// Compare property "altitude"
	if lhs, rhs := this.ActivityStreamsAltitude, o.GetActivityStreamsAltitude(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "attachment"
	if lhs, rhs := this.ActivityStreamsAttachment, o.GetActivityStreamsAttachment(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "attributedTo"
	if lhs, rhs := this.ActivityStreamsAttributedTo, o.GetActivityStreamsAttributedTo(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "audience"
	if lhs, rhs := this.ActivityStreamsAudience, o.GetActivityStreamsAudience(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "bcc"
	if lhs, rhs := this.ActivityStreamsBcc, o.GetActivityStreamsBcc(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "bto"
	if lhs, rhs := this.ActivityStreamsBto, o.GetActivityStreamsBto(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "cc"
	if lhs, rhs := this.ActivityStreamsCc, o.GetActivityStreamsCc(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "content"
	if lhs, rhs := this.ActivityStreamsContent, o.GetActivityStreamsContent(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "context"
	if lhs, rhs := this.ActivityStreamsContext, o.GetActivityStreamsContext(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "current"
	if lhs, rhs := this.ActivityStreamsCurrent, o.GetActivityStreamsCurrent(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "duration"
	if lhs, rhs := this.ActivityStreamsDuration, o.GetActivityStreamsDuration(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "endTime"
	if lhs, rhs := this.ActivityStreamsEndTime, o.GetActivityStreamsEndTime(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "first"
	if lhs, rhs := this.ActivityStreamsFirst, o.GetActivityStreamsFirst(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "generator"
	if lhs, rhs := this.ActivityStreamsGenerator, o.GetActivityStreamsGenerator(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "icon"
	if lhs, rhs := this.ActivityStreamsIcon, o.GetActivityStreamsIcon(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "id"
	if lhs, rhs := this.ActivityStreamsId, o.GetActivityStreamsId(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "image"
	if lhs, rhs := this.ActivityStreamsImage, o.GetActivityStreamsImage(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "inReplyTo"
	if lhs, rhs := this.ActivityStreamsInReplyTo, o.GetActivityStreamsInReplyTo(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "last"
	if lhs, rhs := this.ActivityStreamsLast, o.GetActivityStreamsLast(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "likes"
	if lhs, rhs := this.ActivityStreamsLikes, o.GetActivityStreamsLikes(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "location"
	if lhs, rhs := this.ActivityStreamsLocation, o.GetActivityStreamsLocation(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "mediaType"
	if lhs, rhs := this.ActivityStreamsMediaType, o.GetActivityStreamsMediaType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "name"
	if lhs, rhs := this.ActivityStreamsName, o.GetActivityStreamsName(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "next"
	if lhs, rhs := this.ActivityStreamsNext, o.GetActivityStreamsNext(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "object"
	if lhs, rhs := this.ActivityStreamsObject, o.GetActivityStreamsObject(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "orderedItems"
	if lhs, rhs := this.ActivityStreamsOrderedItems, o.GetActivityStreamsOrderedItems(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "partOf"
	if lhs, rhs := this.ActivityStreamsPartOf, o.GetActivityStreamsPartOf(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "prev"
	if lhs, rhs := this.ActivityStreamsPrev, o.GetActivityStreamsPrev(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "preview"
	if lhs, rhs := this.ActivityStreamsPreview, o.GetActivityStreamsPreview(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "published"
	if lhs, rhs := this.ActivityStreamsPublished, o.GetActivityStreamsPublished(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "replies"
	if lhs, rhs := this.ActivityStreamsReplies, o.GetActivityStreamsReplies(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "shares"
	if lhs, rhs := this.ActivityStreamsShares, o.GetActivityStreamsShares(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "startIndex"
	if lhs, rhs := this.ActivityStreamsStartIndex, o.GetActivityStreamsStartIndex(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "startTime"
	if lhs, rhs := this.ActivityStreamsStartTime, o.GetActivityStreamsStartTime(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "summary"
	if lhs, rhs := this.ActivityStreamsSummary, o.GetActivityStreamsSummary(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "tag"
	if lhs, rhs := this.ActivityStreamsTag, o.GetActivityStreamsTag(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "to"
	if lhs, rhs := this.ActivityStreamsTo, o.GetActivityStreamsTo(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "totalItems"
	if lhs, rhs := this.ActivityStreamsTotalItems, o.GetActivityStreamsTotalItems(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "type"
	if lhs, rhs := this.ActivityStreamsType, o.GetActivityStreamsType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "updated"
	if lhs, rhs := this.ActivityStreamsUpdated, o.GetActivityStreamsUpdated(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "url"
	if lhs, rhs := this.ActivityStreamsUrl, o.GetActivityStreamsUrl(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// End: Compare known properties

	// Begin: Compare unknown properties (only by number of them)
	if len(this.unknown) < len(o.GetUnknownProperties()) {
		return true
	} else if len(o.GetUnknownProperties()) < len(this.unknown) {
		return false
	} // End: Compare unknown properties (only by number of them)

	// All properties are the same.
	return false
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format.
func (this ActivityStreamsOrderedCollectionPage) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	typeName := "OrderedCollectionPage"
	if len(this.alias) > 0 {
		typeName = this.alias + ":" + "OrderedCollectionPage"
	}
	m["type"] = typeName
	// Begin: Serialize known properties
	// Maybe serialize property "altitude"
	if this.ActivityStreamsAltitude != nil {
		if i, err := this.ActivityStreamsAltitude.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAltitude.Name()] = i
		}
	}
	// Maybe serialize property "attachment"
	if this.ActivityStreamsAttachment != nil {
		if i, err := this.ActivityStreamsAttachment.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAttachment.Name()] = i
		}
	}
	// Maybe serialize property "attributedTo"
	if this.ActivityStreamsAttributedTo != nil {
		if i, err := this.ActivityStreamsAttributedTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAttributedTo.Name()] = i
		}
	}
	// Maybe serialize property "audience"
	if this.ActivityStreamsAudience != nil {
		if i, err := this.ActivityStreamsAudience.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAudience.Name()] = i
		}
	}
	// Maybe serialize property "bcc"
	if this.ActivityStreamsBcc != nil {
		if i, err := this.ActivityStreamsBcc.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsBcc.Name()] = i
		}
	}
	// Maybe serialize property "bto"
	if this.ActivityStreamsBto != nil {
		if i, err := this.ActivityStreamsBto.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsBto.Name()] = i
		}
	}
	// Maybe serialize property "cc"
	if this.ActivityStreamsCc != nil {
		if i, err := this.ActivityStreamsCc.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsCc.Name()] = i
		}
	}
	// Maybe serialize property "content"
	if this.ActivityStreamsContent != nil {
		if i, err := this.ActivityStreamsContent.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsContent.Name()] = i
		}
	}
	// Maybe serialize property "context"
	if this.ActivityStreamsContext != nil {
		if i, err := this.ActivityStreamsContext.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsContext.Name()] = i
		}
	}
	// Maybe serialize property "current"
	if this.ActivityStreamsCurrent != nil {
		if i, err := this.ActivityStreamsCurrent.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsCurrent.Name()] = i
		}
	}
	// Maybe serialize property "duration"
	if this.ActivityStreamsDuration != nil {
		if i, err := this.ActivityStreamsDuration.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsDuration.Name()] = i
		}
	}
	// Maybe serialize property "endTime"
	if this.ActivityStreamsEndTime != nil {
		if i, err := this.ActivityStreamsEndTime.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsEndTime.Name()] = i
		}
	}
	// Maybe serialize property "first"
	if this.ActivityStreamsFirst != nil {
		if i, err := this.ActivityStreamsFirst.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsFirst.Name()] = i
		}
	}
	// Maybe serialize property "generator"
	if this.ActivityStreamsGenerator != nil {
		if i, err := this.ActivityStreamsGenerator.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsGenerator.Name()] = i
		}
	}
	// Maybe serialize property "icon"
	if this.ActivityStreamsIcon != nil {
		if i, err := this.ActivityStreamsIcon.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsIcon.Name()] = i
		}
	}
	// Maybe serialize property "id"
	if this.ActivityStreamsId != nil {
		if i, err := this.ActivityStreamsId.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsId.Name()] = i
		}
	}
	// Maybe serialize property "image"
	if this.ActivityStreamsImage != nil {
		if i, err := this.ActivityStreamsImage.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsImage.Name()] = i
		}
	}
	// Maybe serialize property "inReplyTo"
	if this.ActivityStreamsInReplyTo != nil {
		if i, err := this.ActivityStreamsInReplyTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsInReplyTo.Name()] = i
		}
	}
	// Maybe serialize property "last"
	if this.ActivityStreamsLast != nil {
		if i, err := this.ActivityStreamsLast.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsLast.Name()] = i
		}
	}
	// Maybe serialize property "likes"
	if this.ActivityStreamsLikes != nil {
		if i, err := this.ActivityStreamsLikes.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsLikes.Name()] = i
		}
	}
	// Maybe serialize property "location"
	if this.ActivityStreamsLocation != nil {
		if i, err := this.ActivityStreamsLocation.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsLocation.Name()] = i
		}
	}
	// Maybe serialize property "mediaType"
	if this.ActivityStreamsMediaType != nil {
		if i, err := this.ActivityStreamsMediaType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsMediaType.Name()] = i
		}
	}
	// Maybe serialize property "name"
	if this.ActivityStreamsName != nil {
		if i, err := this.ActivityStreamsName.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsName.Name()] = i
		}
	}
	// Maybe serialize property "next"
	if this.ActivityStreamsNext != nil {
		if i, err := this.ActivityStreamsNext.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsNext.Name()] = i
		}
	}
	// Maybe serialize property "object"
	if this.ActivityStreamsObject != nil {
		if i, err := this.ActivityStreamsObject.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsObject.Name()] = i
		}
	}
	// Maybe serialize property "orderedItems"
	if this.ActivityStreamsOrderedItems != nil {
		if i, err := this.ActivityStreamsOrderedItems.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsOrderedItems.Name()] = i
		}
	}
	// Maybe serialize property "partOf"
	if this.ActivityStreamsPartOf != nil {
		if i, err := this.ActivityStreamsPartOf.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPartOf.Name()] = i
		}
	}
	// Maybe serialize property "prev"
	if this.ActivityStreamsPrev != nil {
		if i, err := this.ActivityStreamsPrev.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPrev.Name()] = i
		}
	}
	// Maybe serialize property "preview"
	if this.ActivityStreamsPreview != nil {
		if i, err := this.ActivityStreamsPreview.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPreview.Name()] = i
		}
	}
	// Maybe serialize property "published"
	if this.ActivityStreamsPublished != nil {
		if i, err := this.ActivityStreamsPublished.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPublished.Name()] = i
		}
	}
	// Maybe serialize property "replies"
	if this.ActivityStreamsReplies != nil {
		if i, err := this.ActivityStreamsReplies.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsReplies.Name()] = i
		}
	}
	// Maybe serialize property "shares"
	if this.ActivityStreamsShares != nil {
		if i, err := this.ActivityStreamsShares.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsShares.Name()] = i
		}
	}
	// Maybe serialize property "startIndex"
	if this.ActivityStreamsStartIndex != nil {
		if i, err := this.ActivityStreamsStartIndex.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsStartIndex.Name()] = i
		}
	}
	// Maybe serialize property "startTime"
	if this.ActivityStreamsStartTime != nil {
		if i, err := this.ActivityStreamsStartTime.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsStartTime.Name()] = i
		}
	}
	// Maybe serialize property "summary"
	if this.ActivityStreamsSummary != nil {
		if i, err := this.ActivityStreamsSummary.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsSummary.Name()] = i
		}
	}
	// Maybe serialize property "tag"
	if this.ActivityStreamsTag != nil {
		if i, err := this.ActivityStreamsTag.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsTag.Name()] = i
		}
	}
	// Maybe serialize property "to"
	if this.ActivityStreamsTo != nil {
		if i, err := this.ActivityStreamsTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsTo.Name()] = i
		}
	}
	// Maybe serialize property "totalItems"
	if this.ActivityStreamsTotalItems != nil {
		if i, err := this.ActivityStreamsTotalItems.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsTotalItems.Name()] = i
		}
	}
	// Maybe serialize property "type"
	if this.ActivityStreamsType != nil {
		if i, err := this.ActivityStreamsType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsType.Name()] = i
		}
	}
	// Maybe serialize property "updated"
	if this.ActivityStreamsUpdated != nil {
		if i, err := this.ActivityStreamsUpdated.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsUpdated.Name()] = i
		}
	}
	// Maybe serialize property "url"
	if this.ActivityStreamsUrl != nil {
		if i, err := this.ActivityStreamsUrl.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsUrl.Name()] = i
		}
	}
	// End: Serialize known properties

	// Begin: Serialize unknown properties
	for k, v := range this.unknown {
		// To be safe, ensure we aren't overwriting a known property
		if _, has := m[k]; !has {
			m[k] = v
		}
	}
	// End: Serialize unknown properties

	return m, nil
}

// SetActivityStreamsAltitude returns the "altitude" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsAltitude(i vocab.ActivityStreamsAltitudeProperty) {
	this.ActivityStreamsAltitude = i
}

// SetActivityStreamsAttachment returns the "attachment" property if it exists,
// and nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsAttachment(i vocab.ActivityStreamsAttachmentProperty) {
	this.ActivityStreamsAttachment = i
}

// SetActivityStreamsAttributedTo returns the "attributedTo" property if it
// exists, and nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsAttributedTo(i vocab.ActivityStreamsAttributedToProperty) {
	this.ActivityStreamsAttributedTo = i
}

// SetActivityStreamsAudience returns the "audience" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsAudience(i vocab.ActivityStreamsAudienceProperty) {
	this.ActivityStreamsAudience = i
}

// SetActivityStreamsBcc returns the "bcc" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsBcc(i vocab.ActivityStreamsBccProperty) {
	this.ActivityStreamsBcc = i
}

// SetActivityStreamsBto returns the "bto" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsBto(i vocab.ActivityStreamsBtoProperty) {
	this.ActivityStreamsBto = i
}

// SetActivityStreamsCc returns the "cc" property if it exists, and nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsCc(i vocab.ActivityStreamsCcProperty) {
	this.ActivityStreamsCc = i
}

// SetActivityStreamsContent returns the "content" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsContent(i vocab.ActivityStreamsContentProperty) {
	this.ActivityStreamsContent = i
}

// SetActivityStreamsContext returns the "context" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsContext(i vocab.ActivityStreamsContextProperty) {
	this.ActivityStreamsContext = i
}

// SetActivityStreamsCurrent returns the "current" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsCurrent(i vocab.ActivityStreamsCurrentProperty) {
	this.ActivityStreamsCurrent = i
}

// SetActivityStreamsDuration returns the "duration" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsDuration(i vocab.ActivityStreamsDurationProperty) {
	this.ActivityStreamsDuration = i
}

// SetActivityStreamsEndTime returns the "endTime" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsEndTime(i vocab.ActivityStreamsEndTimeProperty) {
	this.ActivityStreamsEndTime = i
}

// SetActivityStreamsFirst returns the "first" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsFirst(i vocab.ActivityStreamsFirstProperty) {
	this.ActivityStreamsFirst = i
}

// SetActivityStreamsGenerator returns the "generator" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsGenerator(i vocab.ActivityStreamsGeneratorProperty) {
	this.ActivityStreamsGenerator = i
}

// SetActivityStreamsIcon returns the "icon" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsIcon(i vocab.ActivityStreamsIconProperty) {
	this.ActivityStreamsIcon = i
}

// SetActivityStreamsId returns the "id" property if it exists, and nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsId(i vocab.ActivityStreamsIdProperty) {
	this.ActivityStreamsId = i
}

// SetActivityStreamsImage returns the "image" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsImage(i vocab.ActivityStreamsImageProperty) {
	this.ActivityStreamsImage = i
}

// SetActivityStreamsInReplyTo returns the "inReplyTo" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsInReplyTo(i vocab.ActivityStreamsInReplyToProperty) {
	this.ActivityStreamsInReplyTo = i
}

// SetActivityStreamsLast returns the "last" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsLast(i vocab.ActivityStreamsLastProperty) {
	this.ActivityStreamsLast = i
}

// SetActivityStreamsLikes returns the "likes" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsLikes(i vocab.ActivityStreamsLikesProperty) {
	this.ActivityStreamsLikes = i
}

// SetActivityStreamsLocation returns the "location" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsLocation(i vocab.ActivityStreamsLocationProperty) {
	this.ActivityStreamsLocation = i
}

// SetActivityStreamsMediaType returns the "mediaType" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsMediaType(i vocab.ActivityStreamsMediaTypeProperty) {
	this.ActivityStreamsMediaType = i
}

// SetActivityStreamsName returns the "name" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsName(i vocab.ActivityStreamsNameProperty) {
	this.ActivityStreamsName = i
}

// SetActivityStreamsNext returns the "next" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsNext(i vocab.ActivityStreamsNextProperty) {
	this.ActivityStreamsNext = i
}

// SetActivityStreamsObject returns the "object" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsObject(i vocab.ActivityStreamsObjectProperty) {
	this.ActivityStreamsObject = i
}

// SetActivityStreamsOrderedItems returns the "orderedItems" property if it
// exists, and nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsOrderedItems(i vocab.ActivityStreamsOrderedItemsProperty) {
	this.ActivityStreamsOrderedItems = i
}

// SetActivityStreamsPartOf returns the "partOf" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsPartOf(i vocab.ActivityStreamsPartOfProperty) {
	this.ActivityStreamsPartOf = i
}

// SetActivityStreamsPrev returns the "prev" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsPrev(i vocab.ActivityStreamsPrevProperty) {
	this.ActivityStreamsPrev = i
}

// SetActivityStreamsPreview returns the "preview" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsPreview(i vocab.ActivityStreamsPreviewProperty) {
	this.ActivityStreamsPreview = i
}

// SetActivityStreamsPublished returns the "published" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsPublished(i vocab.ActivityStreamsPublishedProperty) {
	this.ActivityStreamsPublished = i
}

// SetActivityStreamsReplies returns the "replies" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsReplies(i vocab.ActivityStreamsRepliesProperty) {
	this.ActivityStreamsReplies = i
}

// SetActivityStreamsShares returns the "shares" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsShares(i vocab.ActivityStreamsSharesProperty) {
	this.ActivityStreamsShares = i
}

// SetActivityStreamsStartIndex returns the "startIndex" property if it exists,
// and nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsStartIndex(i vocab.ActivityStreamsStartIndexProperty) {
	this.ActivityStreamsStartIndex = i
}

// SetActivityStreamsStartTime returns the "startTime" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsStartTime(i vocab.ActivityStreamsStartTimeProperty) {
	this.ActivityStreamsStartTime = i
}

// SetActivityStreamsSummary returns the "summary" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsSummary(i vocab.ActivityStreamsSummaryProperty) {
	this.ActivityStreamsSummary = i
}

// SetActivityStreamsTag returns the "tag" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsTag(i vocab.ActivityStreamsTagProperty) {
	this.ActivityStreamsTag = i
}

// SetActivityStreamsTo returns the "to" property if it exists, and nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsTo(i vocab.ActivityStreamsToProperty) {
	this.ActivityStreamsTo = i
}

// SetActivityStreamsTotalItems returns the "totalItems" property if it exists,
// and nil otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsTotalItems(i vocab.ActivityStreamsTotalItemsProperty) {
	this.ActivityStreamsTotalItems = i
}

// SetActivityStreamsType returns the "type" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsType(i vocab.ActivityStreamsTypeProperty) {
	this.ActivityStreamsType = i
}

// SetActivityStreamsUpdated returns the "updated" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsUpdated(i vocab.ActivityStreamsUpdatedProperty) {
	this.ActivityStreamsUpdated = i
}

// SetActivityStreamsUrl returns the "url" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsOrderedCollectionPage) SetActivityStreamsUrl(i vocab.ActivityStreamsUrlProperty) {
	this.ActivityStreamsUrl = i
}

// VocabularyURI returns the vocabulary's URI as a string.
func (this ActivityStreamsOrderedCollectionPage) VocabularyURI() string {
	return "https://www.w3.org/TR/activitystreams-vocabulary"
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this ActivityStreamsOrderedCollectionPage) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
	if i == nil {
		return toMerge
	}
	for k, v := range i.JSONLDContext() {
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		toMerge[k] = v
	}
	return toMerge
}