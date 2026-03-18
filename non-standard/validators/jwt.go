package validators

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/go-playground/validator/v10"
)

// JWT is the validation function for validating if the current field's value
// is a valid JWT string. Unlike the baked-in jwt validator which only checks
// the format via regex, this performs structural validation: base64url decoding
// and JSON parsing of the header and payload segments.
func JWT(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return false
	}

	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return false
	}

	// Header and payload must be valid base64url-encoded JSON objects
	for _, part := range parts[:2] {
		decoded, err := base64.RawURLEncoding.DecodeString(part)
		if err != nil {
			return false
		}

		var obj map[string]interface{}
		if err := json.Unmarshal(decoded, &obj); err != nil {
			return false
		}
	}

	// Signature segment must be valid base64url (can be empty for unsecured JWTs)
	if parts[2] != "" {
		if _, err := base64.RawURLEncoding.DecodeString(parts[2]); err != nil {
			return false
		}
	}

	return true
}
