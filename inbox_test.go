package rossby_test

import (
	"testing"

	. "github.com/kansaslabs/rossby"
	"github.com/kansaslabs/rossby/bucket"
	"github.com/stretchr/testify/require"
)

func TestBoxID(t *testing.T) {
	box := NewBoxID()

	pbox, err := ParseBoxID(box.String())
	require.NoError(t, err)
	require.Equal(t, pbox, box)

	pbox, err = ParseBoxID(box.Bytes())
	require.NoError(t, err)
	require.Equal(t, pbox, box)

	pbox, err = ParseBoxID(box.Key(bucket.Contacts))
	require.NoError(t, err)
	require.Equal(t, pbox, box)

	pbox, err = ParseBoxID(nil)
	require.NoError(t, err)
	require.Equal(t, pbox, NilBox)

	pbox, err = ParseBoxID(true)
	require.Error(t, err)
	require.Equal(t, pbox, NilBox)
}
