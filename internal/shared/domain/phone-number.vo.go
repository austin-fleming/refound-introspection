package domain

import (
	"errors"
)

type PhoneNumberVO struct {
	value string
}

func ParsePhoneNumber(isRequired bool, value string) (PhoneNumberVO, error) {
	if !isRequired && value == "" {
		return PhoneNumberVO{}, nil
	}

	// TODO: This needs to be validated through a service.
	if value == "" {
		return PhoneNumberVO{}, errors.New("phone number is missing")
	}

	return PhoneNumberVO{value}, nil
}

func (phone PhoneNumberVO) ToValue() string {
	return phone.value
}
