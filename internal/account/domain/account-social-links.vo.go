package domain

import (
	"errors"
	"fmt"
	shared "refound/internal/shared/domain"
)

type AccountSocialLinksVO struct {
	value []string
}

func ParseAccountSocialLinks(isRequired bool, values []string) (AccountSocialLinksVO, error) {
	if !isRequired && len(values) == 0 {
		return AccountSocialLinksVO{}, nil
	}

	var errorIndices []int

	if len(values) == 0 {
		return AccountSocialLinksVO{}, errors.New("account social links are missing")
	}

	var socialLinks AccountSocialLinksVO
	for idx, value := range values {
		if url, err := shared.ParseUrl(true, value); err != nil {
			errorIndices = append(errorIndices, idx)
		} else {
			socialLinks.value = append(socialLinks.value, url.ToValue())
		}
	}

	if len(errorIndices) != 0 {
		return AccountSocialLinksVO{}, fmt.Errorf("social links '%v' are missing", errorIndices)
	}

	return socialLinks, nil
}

func (asl *AccountSocialLinksVO) ToValue() []string {
	return asl.value
}
