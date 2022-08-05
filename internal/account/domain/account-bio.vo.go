package domain

import (
	"errors"
	"strings"
)

type AccountBioVO struct {
	value string
}

func ParseAccountBio(isRequired bool, value string) (AccountBioVO, error) {
	if !isRequired && value == "" {
		return AccountBioVO{}, nil
	}

	// TODO: should validate as markdown
	if strings.TrimSpace(value) == "" {
		return AccountBioVO{}, errors.New("account bio is missing")
	}

	return AccountBioVO{value}, nil
}

func (bio *AccountBioVO) ToValue() string {
	return bio.value
}
