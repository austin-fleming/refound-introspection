package validation

import "testing"

func TestUuidV4(test *testing.T) {
	cases := []struct {
		arg  string
		want bool
	}{
		{"57b73598-8764-4ad0-a76a-679bb6640eb1", true},
		{"625e63f3-58f5-40b7-83a1-a72ad31acffb", true},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-5078-af07-9141ba07c9f3", false},
		{"934859", false},
		{"", false},
	}

	for _, tc := range cases {
		got := IsUuidV4(tc.arg)
		if tc.want != got {
			test.Errorf("Expected '%t' for '%s', but got '%t'", tc.want, tc.arg, got)
		}
	}
}
