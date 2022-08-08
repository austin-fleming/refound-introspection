package domain

import (
	"fmt"
	validation "refound/internal/shared/utils/validation"
)

type UrlVO struct {
	value string
}

func ParseUrl(isRequired bool, value string) (UrlVO, error) {
	if !isRequired && value == "" {
		return UrlVO{}, nil
	}

	if validation.IsSecureUrl(value) {
		return UrlVO{}, fmt.Errorf("%s is not a valid secure url", value)
	}

	return UrlVO{value}, nil
}

func (u UrlVO) ToValue() string {
	return u.value
}
