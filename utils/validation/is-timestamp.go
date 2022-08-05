package validation

import "time"

const (
	rfc3339WithoutZone = "2006-01-02T15:04:05"
)

func IsTime(str string, format string) bool {
	_, err := time.Parse(format, str)
	return err == nil
}

func IsTimeRFC3339(str string) bool {
	return IsTime(str, time.RFC3339) || IsTime(str, rfc3339WithoutZone)
}
