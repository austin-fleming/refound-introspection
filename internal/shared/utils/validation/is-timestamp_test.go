package validation

import "testing"

func TestIsTimeRFC3339(test *testing.T) {
	// test cases from https://github.com/asaskevich/govalidator/blob/f21760c49a8d602d863493de796926d2a5c1138d/validator_test.go
	testCases := []struct {
		time string
		want bool
	}{
		{"2016-12-31 11:00", false},
		{"2016-12-31 11:00:00", false},
		{"2016-12-31T11:00", false},
		{"2016-12-31T11:00:00", true},
		{"2016-12-31T11:00:00Z", true},
		{"2016-12-31T11:00:00+01:00", true},
		{"2016-12-31T11:00:00-01:00", true},
		{"2016-12-31T11:00:00.05Z", true},
		{"2016-12-31T11:00:00.05-01:00", true},
		{"2016-12-31T11:00:00.05+01:00", true},
		{"2016-12-31T11:00:00", true},
		{"2016-12-31T11:00:00Z", true},
		{"2016-12-31T11:00:00+01:00", true},
		{"2016-12-31T11:00:00-01:00", true},
		{"2016-12-31T11:00:00.05Z", true},
		{"2016-12-31T11:00:00.05-01:00", true},
		{"2016-12-31T11:00:00.05+01:00", true},
	}

	for _, tc := range testCases {
		got := IsTimeRFC3339(tc.time)
		if tc.want != got {
			test.Errorf("\nExpected:\n%t\nGot:\n%t", tc.want, got)
		}
	}
}
