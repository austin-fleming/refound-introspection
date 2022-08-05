package domain

import (
	"fmt"
	"strings"

	v4 "github.com/google/uuid"
)

type AccountNonceVO struct {
	value string
}

func ParseAccountNonce(isRequired bool, value string) (AccountNonceVO, error) {
	if !isRequired && value == "" {
		return AccountNonceVO{}, nil
	}

	preparedValue := strings.TrimSpace(value)

	if preparedValue == "" {
		return AccountNonceVO{}, fmt.Errorf("nonce is missing")
	}

	return AccountNonceVO{value: preparedValue}, nil
}

func GenerateAccountNonce() AccountNonceVO {
	return AccountNonceVO{value: v4.New().String()}
}

func (nonce AccountNonceVO) ToValue() string {
	return nonce.value
}

func (nonce AccountNonceVO) Regenerate() AccountNonceVO {
	return GenerateAccountNonce()
}
