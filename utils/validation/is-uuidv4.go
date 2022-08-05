package validation

import "regexp"

var rxUuidV4 = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")

func IsUuidV4(str string) bool {
	return rxUuidV4.MatchString(str)
}
