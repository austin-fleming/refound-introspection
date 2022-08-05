package domain

import (
	"fmt"
	"regexp"
)

type UsernameVO struct {
	value string
}

var rxUsername = regexp.MustCompile("^[a-zA-Z]([a-zA-Z0-9]{0,28}|[a-zA-Z0-9]{0,26}[_.]?[a-zA-Z0-9]{0,26}|[a-zA-Z0-9]{0,27}[_.]?|[_.]?[a-zA-Z0-9]{0,27})[a-zA-Z0-9]$")

func ParseUsername(isRequired bool, value string) (UsernameVO, error) {
	if !isRequired && value == "" {
		return UsernameVO{}, nil
	}

	if rxUsername.MatchString(value) {
		return UsernameVO{}, fmt.Errorf("%s is not a valid username", value)
	}

	return UsernameVO{value}, nil
}

func (u UsernameVO) ToValue() string {
	return u.value
}
