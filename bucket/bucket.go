/*
Package bucket implements prefixed-keys to separate collections of similar data in the
database. While this package could be made database-agnostic, it is intended to optimize
scans and searches over BadgerDB. BadgerDB (and many other embedded key-value databases)
do not provide collections or buckets to group related keys. This package allows us to
prefix a key to achieve similar functionality and still maintain cross-bucket
transactions inside of the same database. The prefixes are fixed length to ensure that
they can be easily parsed from the key.

This package is separate from other packages to provide easy one line functionality, to
access buckets e.g. bucket.Profiles.
*/
package bucket

// Bucket prefix constants for named lookup.
const (
	Unassigned Bucket = iota
	PublicKeys
	Messages
	Contacts
	Profiles
)

// String values of bucket names for human readability.
var bucketNames = []string{
	"unassigned",
	"public keys",
	"messages",
	"contacts",
	"profiles",
}

// Bucket implements a one byte prefix that can be added to a key. Note that the use of
// uint8 limits us to only 255 unique buckets. If we require more buckets we should
// ensure that the size of the bucket prefix remains fixed (e.g. uint16).
type Bucket uint8

// Key is a bucket prefixed byte array that can be used to Get and Set values in the
// database. Keys can also be parsed to retreive the key and prefixe components.
type Key []byte

// String returns the human readable name of the bucket, not the actual prefix.
func (b Bucket) String() string {
	if int(b) < len(bucketNames) {
		return bucketNames[b]
	}
	return "unknown bucket"
}

// Key creates a bucket prefixed key and returns it.
func (b Bucket) Key(key []byte) Key {
	return append([]byte{uint8(b)}, key...)
}

// Parse a key into its bucket and original key values.
func (k Key) Parse() (Bucket, []byte) {
	return Bucket(k[0]), k[1:]
}
