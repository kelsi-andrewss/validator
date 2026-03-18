package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var jwtRegex = regexp.MustCompile(`^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]*$`)

// IsJWT is the validation function for validating if the current field
// is a valid JSON Web Token (RFC 7519).
func IsJWT(fl validator.FieldLevel) bool {
	return jwtRegex.MatchString(fl.Field().String())
}
