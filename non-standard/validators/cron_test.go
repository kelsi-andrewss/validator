package validators

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func TestIsCron(t *testing.T) {
	tests := []struct {
		value string
		valid bool
	}{
		// Valid - standard 5-field cron
		{"* * * * *", true},
		{"0 0 * * *", true},
		{"0 0 1 1 *", true},
		{"*/5 * * * *", true},
		{"0 0 1,15 * *", true},
		{"0 0 1-5 * *", true},
		{"5 4 * * 0", true},
		{"30 18 * * 1,2,3,4,5", true},

		// Valid - 6/7 field cron (with seconds / year)
		{"0 0 0 * * *", true},
		{"0 0 0 * * * *", true},

		// Valid - predefined schedules
		{"@annually", true},
		{"@yearly", true},
		{"@monthly", true},
		{"@weekly", true},
		{"@daily", true},
		{"@hourly", true},
		{"@reboot", true},

		// Valid - interval
		{"@every 5m", true},
		{"@every 1h30m", true},
		{"@every 500ms", true},
		{"@every 10s", true},

		// Invalid
		{"", false},
		{"not a cron", false},
		{"* * *", false},
		{"* * * *", false},
	}

	v := validator.New()
	err := v.RegisterValidation("cron", IsCron)
	assert.Equal(t, nil, err)

	type field struct {
		Cron string `validate:"cron"`
	}

	for _, tc := range tests {
		err := v.Struct(field{Cron: tc.value})
		if tc.valid {
			assert.Equal(t, nil, err)
		} else {
			assert.NotEqual(t, nil, err)
		}
	}
}
