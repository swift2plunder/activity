package vocab

import (
	"net/url"
	"time"
)

// The date and time describing the actual or expected ending time of the object.
// When used with an Activity object, for instance, the endTime property
// specifies the moment the activity concluded or is expected to conclude.
//
// Example 127 (https://www.w3.org/TR/activitystreams-vocabulary/#ex144-jsonld):
//   {
//     "endTime": "2015-01-01T06:00:00-08:00",
//     "name": "Going-Away Party for Jim",
//     "startTime": "2014-12-31T23:00:00-08:00",
//     "type": "Event"
//   }
type EndTimePropertyInterface interface {
	// Clear ensures no value of this property is set. Calling IsDateTime
	// afterwards will return false.
	Clear()
	// Get returns the value of this property. When IsDateTime returns false,
	// Get will return any arbitrary value.
	Get() time.Time
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsDateTime returns true if this property is set and not an IRI.
	IsDateTime() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o EndTimePropertyInterface) bool
	// Name returns the name of this property: "endTime".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets the value of this property. Calling IsDateTime afterwards will
	// return true.
	Set(v time.Time)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}