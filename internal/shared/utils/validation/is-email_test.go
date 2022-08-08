package validation

import "testing"

func TestIsEmail(test *testing.T) {
	cases := []struct {
		arg  string
		want bool
	}{
		{"name@example.com", true},
		{"name@.com", false},
		{"name@example", false},
		{"nameexample.com", false},
		{"", false},
	}

	for _, tc := range cases {
		got := IsEmail(tc.arg)
		if tc.want != got {
			test.Errorf("Expected '%t', but got '%t'", tc.want, got)
		}
	}
}
