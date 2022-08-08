package domain

import (
	"fmt"
	validation "refound/internal/shared/utils/validation"

	v4 "github.com/google/uuid"
)

type IdVO struct {
	value string
}

func validate(value string) error {
	if !validation.IsUuidV4(value) {
		return fmt.Errorf("%s is not a valid Id", value)
	}

	return nil
}

func ParseId(isRequired bool, value string) (IdVO, error) {
	if !isRequired && value == "" {
		return IdVO{}, nil
	}
	if err := validate(value); err != nil {
		return IdVO{}, err
	}

	return IdVO{value}, nil
}

func GenerateId() IdVO {
	return IdVO{value: v4.New().String()}
}

func (id IdVO) ToValue() string {
	return id.value
}
