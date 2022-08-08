package domain

import (
	"fmt"
	validation "refound/internal/shared/utils/validation"
)

type EthAddressVO struct {
	value string
}

func ParseEthAddress(isRequired bool, value string) (EthAddressVO, error) {
	if !isRequired && value == "" {
		return EthAddressVO{}, nil
	}

	if !validation.IsEthAddress(value) {
		return EthAddressVO{}, fmt.Errorf("%s is not a valid eth address", value)
	}

	return EthAddressVO{value}, nil
}

func (ea EthAddressVO) ToValue() string {
	return ea.value
}
