package datetime

import (
	"fmt"
	"time"
)

// SerializeDateTime converts a dateTime value to an interface representation
// suitable for marshalling into a text or binary format.
func SerializeDateTime(this time.Time) (interface{}, error) {
	return this.Format(time.RFC3339), nil
}

// DeserializeDateTime creates dateTime value from an interface representation
// that has been unmarshalled from a text or binary format.
func DeserializeDateTime(this interface{}) (time.Time, error) {
	var tmp time.Time
	var err error
	if s, ok := this.(string); ok {
		tmp, err = time.Parse(time.RFC3339, s)
		if err != nil {
			tmp, err = time.Parse("2006-01-02T15:04Z07:00", s)
			if err != nil {
				err = fmt.Errorf("%v cannot be interpreted as xsd:datetime", this)
			}
		}
	} else {
		err = fmt.Errorf("%v cannot be interpreted as a string for xsd:datetime", this)
	}
	return tmp, err
}

// LessDateTime returns true if the left dateTime value is less than the right
// value.
func LessDateTime(lhs, rhs time.Time) bool {
	return lhs.Before(rhs)
}
