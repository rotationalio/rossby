package pb_test

import (
	"testing"

	. "github.com/kansaslabs/rossby/pb"
	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	errs := Errors{
		Errorf(1, "something bad happened"),
		Errorf(2, "it happened in the %s", "solarium"),
		Errorf(3, "it happened %d times", 23),
	}

	require.Len(t, errs.Serialize().Errors, len(errs))
}
