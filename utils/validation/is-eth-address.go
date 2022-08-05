package validation

import "regexp"

var rxEthAddress = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func IsEthAddress(str string) bool {
	return rxEthAddress.MatchString(str)
}
