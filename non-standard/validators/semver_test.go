package validators

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func TestIsSemver(t *testing.T) {
	tests := []struct {
		value string
		valid bool
	}{
		// Valid
		{"0.0.0", true},
		{"1.0.0", true},
		{"0.1.0", true},
		{"0.0.1", true},
		{"1.2.3", true},
		{"10.20.30", true},
		{"1.0.0-alpha", true},
		{"1.0.0-alpha.1", true},
		{"1.0.0-0.3.7", true},
		{"1.0.0-x.7.z.92", true},
		{"1.0.0+build.1", true},
		{"1.0.0-beta+exp.sha.5114f85", true},
		{"1.0.0+20130313144700", true},
		{"1.0.0-alpha+001", true},
		{"1.2.3-rc.1+build.123", true},
		{"999.999.999", true},

		// Invalid
		{"", false},
		{"1", false},
		{"1.2", false},
		{"v1.2.3", false},
		{"1.2.3.4", false},
		{"01.0.0", false},
		{"1.02.0", false},
		{"1.0.03", false},
		{"1.0.0-", false},
		{"abc", false},
		{"1.0.0-beta!", false},
	}

	v := validator.New()
	err := v.RegisterValidation("semver", IsSemver)
	assert.Equal(t, nil, err)

	type field struct {
		Semver string `validate:"semver"`
	}

	for _, tc := range tests {
		err := v.Struct(field{Semver: tc.value})
		if tc.valid {
			assert.Equal(t, nil, err)
		} else {
			assert.NotEqual(t, nil, err)
		}
	}
}
