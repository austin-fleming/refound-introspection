package validation

import "net/url"

func IsSecureUrl(str string) bool {
	parsed, err := url.ParseRequestURI(str)

	if err != nil {
		return false
	}

	if parsed.Scheme != "https" {
		return false
	}

	return true
}
