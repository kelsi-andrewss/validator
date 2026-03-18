package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var cronRegex = regexp.MustCompile(`^` +
	`(@(annually|yearly|monthly|weekly|daily|hourly|reboot))` +
	`|(@every (\d+(ns|us|µs|ms|s|m|h))+)` +
	`|(((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7}` +
	`$`)

// IsCron is the validation function for validating if the current field
// is a valid cron expression.
func IsCron(fl validator.FieldLevel) bool {
	return cronRegex.MatchString(fl.Field().String())
}
