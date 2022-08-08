package validation

import "testing"

func TestEthAddress(test *testing.T) {
	cases := []struct {
		arg  string
		want bool
	}{
		{"0x52908400098527886E0F7030069857D2E4169EE7", true},
		{"0xde709f2102306220921060314715629080e2fb77", true},
		{"0x02F9AE5f22EA3fA88F05780B30385bECFacbf130", true},
		{"02F9AE5f22EA3fA88F05780B30385bECFacbf130", false},    // no prefix
		{"0xde709f2102306220921060314715629080e2fb7", false},   // too short
		{"0xde709f2102306220921060314715629080e2fb777", false}, // too long
		{"0xde709f2102306220921060314715629080e2fb7g", false},  // non-hex character
		{"", false},
	}

	for _, tc := range cases {
		got := IsEthAddress(tc.arg)
		if tc.want != got {
			test.Errorf("Expected '%t' for '%s', but got '%t'", tc.want, tc.arg, got)
		}
	}
}
