package validators

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var cronFieldRegex = regexp.MustCompile(`^(\*|[0-9]+(-[0-9]+)?)(/[0-9]+)?$`)

// Cron is the validation function for validating if the current field's
// value is a valid standard 5-field cron expression (minute hour dom month dow).
func Cron(fl validator.FieldLevel) bool {
	s := strings.TrimSpace(fl.Field().String())
	if s == "" {
		return false
	}

	fields := strings.Fields(s)
	if len(fields) != 5 {
		return false
	}

	limits := [5][2]int{
		{0, 59},  // minute
		{0, 23},  // hour
		{1, 31},  // day of month
		{1, 12},  // month
		{0, 7},   // day of week (0 and 7 are both Sunday)
	}

	for i, field := range fields {
		if !validateCronField(field, limits[i][0], limits[i][1]) {
			return false
		}
	}

	return true
}

func validateCronField(field string, min, max int) bool {
	for _, part := range strings.Split(field, ",") {
		if !cronFieldRegex.MatchString(part) {
			return false
		}

		// Strip step suffix for range/value validation
		base := part
		if idx := strings.Index(part, "/"); idx != -1 {
			base = part[:idx]
			step, err := strconv.Atoi(part[idx+1:])
			if err != nil || step < 1 {
				return false
			}
		}

		if base == "*" {
			continue
		}

		if dashIdx := strings.Index(base, "-"); dashIdx != -1 {
			lo, err := strconv.Atoi(base[:dashIdx])
			if err != nil || lo < min || lo > max {
				return false
			}
			hi, err := strconv.Atoi(base[dashIdx+1:])
			if err != nil || hi < min || hi > max || hi < lo {
				return false
			}
		} else {
			val, err := strconv.Atoi(base)
			if err != nil || val < min || val > max {
				return false
			}
		}
	}

	return true
}
