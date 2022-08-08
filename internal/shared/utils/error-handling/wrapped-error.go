package errorhandling

type WrappedError struct {
	Detail     string `json:"detail,omitempty`
	Op         bool   `json:"op,omitempty`
	InnerError error  `json:innerError,omitempty`
}



