package rossby

import (
	"fmt"

	"github.com/kansaslabs/rossby/bucket"
	"github.com/segmentio/ksuid"
)

// NilBox is the zero-valued box
var NilBox = BoxID(ksuid.Nil)

// BoxID implements a globally unique distributed identifier for inboxes.
type BoxID ksuid.KSUID

// NewBoxID creates a unique ID based on the ksuid globally unique identifier framework.
func NewBoxID() BoxID {
	return BoxID(ksuid.New())
}

// ParseBoxID from a string, bytes, Key, or nil
func ParseBoxID(obj interface{}) (box BoxID, err error) {
	var uuid ksuid.KSUID

	switch v := obj.(type) {
	case []byte:
		uuid, err = ksuid.FromBytes(v)
	case bucket.Key:
		_, data := v.Parse()
		uuid, err = ksuid.FromBytes(data)
	case string:
		uuid, err = ksuid.Parse(v)
	case nil:
		return NilBox, nil
	default:
		return NilBox, fmt.Errorf("could not parse BoxID from type %s", v)
	}

	if err != nil {
		return NilBox, err
	}
	return BoxID(uuid), nil
}

// Key encodes the BoxID for bucket access
func (id BoxID) Key(b bucket.Bucket) bucket.Key {
	return b.Key(ksuid.KSUID(id).Bytes())
}

// String returns the string encoded BoxID
func (id BoxID) String() string {
	return ksuid.KSUID(id).String()
}

// Bytes returns the raw bytes of the underlying BoxID
func (id BoxID) Bytes() []byte {
	return ksuid.KSUID(id).Bytes()
}
