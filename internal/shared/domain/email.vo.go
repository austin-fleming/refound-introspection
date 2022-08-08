package domain

import (
	"fmt"
	validation "refound/internal/shared/utils/validation"
)

type EmailVO struct {
	value string
}

func ParseEmail(isRequired bool, value string) (EmailVO, error) {
	if !isRequired && value == "" {
		return EmailVO{}, nil
	}

	if !validation.IsEmail(value) {
		return EmailVO{}, fmt.Errorf("%s is not an email", value)
	}

	return EmailVO{value}, nil
}

func (email EmailVO) ToValue() string {
	return email.value
}
