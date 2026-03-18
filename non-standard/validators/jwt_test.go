package validators

import (
	"encoding/base64"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func TestJWT(t *testing.T) {
	type jwtTest struct {
		Token string `validate:"jwt"`
	}

	v := validator.New()
	err := v.RegisterValidation("jwt", JWT)
	assert.Equal(t, nil, err)

	encode := func(s string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(s))
	}

	validHeader := encode(`{"alg":"HS256","typ":"JWT"}`)
	validPayload := encode(`{"sub":"1234567890","name":"Test"}`)
	validSig := base64.RawURLEncoding.EncodeToString([]byte("signature"))

	tests := []struct {
		name    string
		value   string
		isValid bool
	}{
		// Valid
		{"standard JWT", validHeader + "." + validPayload + "." + validSig, true},
		{"unsecured JWT (empty sig)", validHeader + "." + validPayload + ".", true},
		{"minimal header", encode(`{"alg":"none"}`) + "." + encode(`{}`) + ".", true},

		// Invalid
		{"empty string", "", false},
		{"two segments", validHeader + "." + validPayload, false},
		{"one segment", validHeader, false},
		{"four segments", validHeader + "." + validPayload + "." + validSig + ".extra", false},
		{"invalid base64 header", "!!!." + validPayload + "." + validSig, false},
		{"invalid base64 payload", validHeader + ".!!!." + validSig, false},
		{"header not JSON object", encode("not json") + "." + validPayload + "." + validSig, false},
		{"payload not JSON object", validHeader + "." + encode("not json") + "." + validSig, false},
		{"header is JSON array", encode(`[1,2,3]`) + "." + validPayload + "." + validSig, false},
		{"invalid base64 signature", validHeader + "." + validPayload + ".===badpad===", false},
		{"just dots", "..", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := jwtTest{Token: tc.value}
			err := v.Struct(s)
			if tc.isValid {
				assert.Equal(t, nil, err)
			} else {
				assert.NotEqual(t, nil, err)
			}
		})
	}
}
