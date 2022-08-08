package errorhandling

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	// ContentTypeJSON https://tools.ietf.org/html/rfc7807#section-6.1
	ContentTypeJSON = "application/problem+json"
)

/*
Code should be a standard http error code, as opposed to internal codes.
Error Title should be gleaned from the Code, not explicitly stated.
*/
const (
	EInvalid   = fiber.StatusBadRequest
	ENotFound  = fiber.StatusNotFound
	EConflict  = fiber.StatusConflict
	ERateLimit = fiber.StatusTooManyRequests
	EInternal  = fiber.StatusInternalServerError
)

var ETitles = map[int]string{
	400: "bad request",
	404: "not found",
	409: "conflict",
	429: "too many requests",
	500: "internal server error",
}

type ApiError struct {
	Code          int            `json:"status"`
	Detail        string         `json:"detail,omitempty"`
	InvalidFields []InvalidField `json:"invalidFields,omitempty"`
	Op            string         `json:"op,omitempty"`
	Internal      string         `json:"internal,omitempty"`
	Inner         error          `json:"inner,omitempty"`
}

// If standard error, return Internal error code.
// If ApiError, either get the code or recurse down through inner errors.
func ErrorCode(err error) int {
	e, ok := err.(ApiError)

	if ok && e.Code != 0 {
		return e.Code
	}

	if ok && e.Inner != nil {
		return ErrorCode(e.Inner)
	}

	return EInternal
}

func ErrorDetail(err error) string {
	e, ok := err.(ApiError)

	if ok && e.Detail != "" {
		return e.Detail
	}

	if ok && e.Inner != nil {
		return ErrorDetail(e.Inner)
	}

	return ""
}

func ErrorInvalidFields(err error) []InvalidField {
	e, ok := err.(ApiError)

	if ok && len(e.InvalidFields) > 0 {
		return e.InvalidFields
	}

	if ok && e.Inner != nil {
		return ErrorInvalidFields(e.Inner)
	}

	return []InvalidField{}
}

func PrepareError(err error) ApiError {
	e, ok := err.(ApiError)

	// standard errors
	if !ok {
		return ApiError{
			Code:   EInternal,
			Detail: "unexpected error",
			Inner:  err,
		}
	}

	return ApiError{
		Code:          ErrorCode(e),
		Detail:        ErrorDetail(e),
		InvalidFields: ErrorInvalidFields(e),
		Inner:         e,
	}
}

func ErrorToByte(err error) []byte {
	preparedError := PrepareError(err)

	out, err := json.Marshal(preparedError)
	if err != nil {
		// TODO: don't use panic.
		panic(fmt.Sprintf("Something went wrong marshalling '%#v'", err))
	}

	return out
}

func ErrorToString(err error) string {
	return string(ErrorToByte(err))
}

func Error(err error) string {
	return ErrorToString(err)
}

func (err ApiError) Error() string {
	return Error(err)
}

func (err *ApiError) AppendInvalidField(invalidField InvalidField) {
	if (invalidField == InvalidField{}) {
		return
	}
	err.InvalidFields = append(err.InvalidFields, invalidField)
}

// Takes in an error. If error is an ApiError, it's InvalidFields are appended.
// Provided variable number of paths, the paths will be prepended to the added InvalidFields.
func (err *ApiError) AppendInvalidFieldsFromError(otherError error, paths ...string) {
	apiErr, ok := otherError.(ApiError)
	if !ok {
		return
	}

	// Add paths
	if len(paths) > 0 {
		for idx := range apiErr.InvalidFields {
			apiErr.InvalidFields[idx].PrependPath(paths...)
		}
	}

	err.InvalidFields = append(err.InvalidFields, apiErr.InvalidFields...)
}
