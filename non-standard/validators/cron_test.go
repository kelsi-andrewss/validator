package validators

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func TestCron(t *testing.T) {
	type cronTest struct {
		Expression string `validate:"cron"`
	}

	v := validator.New()
	err := v.RegisterValidation("cron", Cron)
	assert.Equal(t, nil, err)

	tests := []struct {
		name    string
		value   string
		isValid bool
	}{
		// Valid
		{"every minute", "* * * * *", true},
		{"specific minute", "0 * * * *", true},
		{"step minutes", "*/5 * * * *", true},
		{"midnight daily", "0 0 * * *", true},
		{"new year midnight", "0 0 1 1 0", true},
		{"range in minutes", "1-30 * * * *", true},
		{"comma separated", "0,15,30,45 * * * *", true},
		{"complex expression", "*/10 9-17 * * 1-5", true},
		{"day of week sunday 7", "0 0 * * 7", true},
		{"all max values", "59 23 31 12 7", true},

		// Invalid
		{"empty string", "", false},
		{"minute too high", "60 * * * *", false},
		{"hour too high", "0 24 * * *", false},
		{"dom too high", "0 0 32 * *", false},
		{"month too high", "0 0 * 13 *", false},
		{"dow too high", "0 0 * * 8", false},
		{"dom zero", "0 0 0 * *", false},
		{"month zero", "0 0 * 0 *", false},
		{"too few fields", "* * * *", false},
		{"too many fields", "* * * * * *", false},
		{"text value", "bad", false},
		{"letters in field", "a * * * *", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := cronTest{Expression: tc.value}
			err := v.Struct(s)
			if tc.isValid {
				assert.Equal(t, nil, err)
			} else {
				assert.NotEqual(t, nil, err)
			}
		})
	}
}
