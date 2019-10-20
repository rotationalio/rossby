package bucket_test

import (
	"testing"

	. "github.com/kansaslabs/rossby/bucket"
	"github.com/stretchr/testify/require"
)

func TestBucketString(t *testing.T) {
	require.Equal(t, Unassigned.String(), "unassigned")
	require.Equal(t, Messages.String(), "messages")
	require.Equal(t, Bucket(255).String(), "unknown bucket")
}

func TestBucketKey(t *testing.T) {
	key := []byte("foo")
	pkey := Profiles.Key(key)

	require.Len(t, pkey, len(key)+1)

	b, k := pkey.Parse()
	require.Equal(t, key, k)
	require.Equal(t, b, Profiles)
}
