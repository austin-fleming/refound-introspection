package domain

import (
	"fmt"
	validation "refound/internal/shared/utils/validation"
	"time"
)

type TimestampVO struct {
	value string
}

func ParseTimestamp(isRequired bool, value string) (TimestampVO, error) {
	if !isRequired && value == "" {
		return TimestampVO{}, nil
	}

	if validation.IsTimeRFC3339(value) {
		return TimestampVO{}, fmt.Errorf("")
	}

	return TimestampVO{value}, nil
}

func GenerateTimestamp() TimestampVO {
	now := time.Now().Format("2006-01-02T15:04:05")
	return TimestampVO{value: now}
}

func (timestamp TimestampVO) ToValue() string {
	return timestamp.value
}
