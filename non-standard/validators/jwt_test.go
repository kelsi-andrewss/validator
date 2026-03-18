package validators

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func TestIsJWT(t *testing.T) {
	tests := []struct {
		value string
		valid bool
	}{
		// Valid - standard 3-part tokens
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U", true},
		{"a.b.c", true},
		{"header.payload.signature", true},
		{"abc-def_ghi.jkl-mno_pqr.stu-vwx_yz", true},
		{"a.b.", true}, // empty signature is valid per regex (signature segment allows zero-length)

		// Invalid
		{"", false},
		{"abc", false},
		{"a.b", false},
		{"a.b.c.d", false},
		{"..", false},
		{"a..c", false},
		{"header.pay load.sig", false},
		{"header.payload.sig nature", false},
	}

	v := validator.New()
	err := v.RegisterValidation("jwt", IsJWT)
	assert.Equal(t, nil, err)

	type field struct {
		JWT string `validate:"jwt"`
	}

	for _, tc := range tests {
		err := v.Struct(field{JWT: tc.value})
		if tc.valid {
			assert.Equal(t, nil, err)
		} else {
			assert.NotEqual(t, nil, err)
		}
	}
}
