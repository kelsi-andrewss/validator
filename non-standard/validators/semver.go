package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// semverRegex validates semantic versioning strings per https://semver.org/
// Supports: major.minor.patch with optional pre-release and build metadata
// Examples: 1.0.0, 1.2.3-alpha.1, 1.0.0-beta+build.123
var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)` +
	`(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?` +
	`(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// Semver is the validation function for validating if the current field's
// value is a valid semantic version string.
func Semver(fl validator.FieldLevel) bool {
	return semverRegex.MatchString(fl.Field().String())
}
