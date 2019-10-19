package pb_test

import (
	"testing"

	. "github.com/kansaslabs/rossby/pb"
	"github.com/stretchr/testify/require"
)

func TestContactValidate(t *testing.T) {
	var phonetests = []struct {
		contact string
		valid   bool
	}{
		{"+14155552671", true},
		{"+442071838750", true},
		{"+551155256325", true},
		{"+1", false},
		{"+55115525632532335", false},
		{"foo", false},
	}

	for _, tt := range phonetests {
		t.Run(tt.contact, func(t *testing.T) {
			contact := Contact{Type: ContactType_PHONE, Contact: tt.contact}
			if tt.valid {
				require.NoError(t, contact.Validate())
			} else {
				require.Error(t, contact.Validate())
			}
		})
	}

	var emailtests = []struct {
		contact string
		valid   bool
	}{
		{"joe@example.com", true},
		{"jane.doe@example.co.uk", true},
		{"foo", false},
		{"+155555555455", false},
	}

	for _, tt := range emailtests {
		t.Run(tt.contact, func(t *testing.T) {
			contact := Contact{Type: ContactType_EMAIL, Contact: tt.contact}
			if tt.valid {
				require.NoError(t, contact.Validate())
			} else {
				require.Error(t, contact.Validate())
			}
		})
	}
}
