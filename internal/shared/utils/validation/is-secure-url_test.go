package validation

import "testing"

func TestIsSecureUrl(test *testing.T) {
	testCases := []struct {
		arg  string
		want bool
	}{
		{
			`https://example.com`,
			true,
		},
		{
			`https://example.com/status/help?var=test`,
			true,
		},
		{
			`http://example.com`,
			false,
		},
		{
			`example.com`,
			false,
		},
	}

	for _, tc := range testCases {
		got := IsSecureUrl(tc.arg)
		if tc.want != got {
			test.Errorf("\nExpected:\n\t%t\nGot:\n\t%t", tc.want, got)
		}
	}
}
