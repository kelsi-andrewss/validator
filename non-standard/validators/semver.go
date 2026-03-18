package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)` +
	`(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?` +
	`(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// IsSemver is the validation function for validating if the current field
// is a valid semantic version string (https://semver.org/).
func IsSemver(fl validator.FieldLevel) bool {
	return semverRegex.MatchString(fl.Field().String())
}
