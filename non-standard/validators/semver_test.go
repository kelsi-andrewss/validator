package validators

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func TestSemver(t *testing.T) {
	type semverTest struct {
		Version string `validate:"semver"`
	}

	v := validator.New()
	err := v.RegisterValidation("semver", Semver)
	assert.Equal(t, nil, err)

	tests := []struct {
		name    string
		value   string
		isValid bool
	}{
		// Valid
		{"basic version", "1.2.3", true},
		{"all zeros", "0.0.0", true},
		{"zero patch", "0.0.1", true},
		{"large numbers", "100.200.300", true},
		{"pre-release alpha", "1.0.0-alpha", true},
		{"pre-release dotted", "1.0.0-alpha.1", true},
		{"pre-release beta", "1.0.0-0.3.7", true},
		{"build metadata", "1.0.0+build.123", true},
		{"pre-release and build", "1.0.0-alpha+001", true},
		{"complex pre-release", "1.0.0-alpha.1+build.123", true},

		// Invalid
		{"empty string", "", false},
		{"two segments", "1.2", false},
		{"v prefix", "v1.2.3", false},
		{"four segments", "1.2.3.4", false},
		{"leading zero major", "01.2.3", false},
		{"leading zero minor", "1.02.3", false},
		{"leading zero patch", "1.2.03", false},
		{"negative major", "-1.2.3", false},
		{"just text", "not-a-version", false},
		{"missing patch", "1.0", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := semverTest{Version: tc.value}
			err := v.Struct(s)
			if tc.isValid {
				assert.Equal(t, nil, err)
			} else {
				assert.NotEqual(t, nil, err)
			}
		})
	}
}
