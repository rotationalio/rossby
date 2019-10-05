package rossby_test

import (
	"testing"

	. "github.com/kansaslabs/rossby"
	"github.com/stretchr/testify/require"
)

const expectedVersion = "0.0"

// TestVersion is a sanity check to ensure that package version upgrades are correct.
func TestVersion(t *testing.T) {
	require.Equal(t, Version(false), expectedVersion)
}
