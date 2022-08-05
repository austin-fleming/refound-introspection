package account_service

type AccountExistsRequestDTO struct {
	Id string `json:"id"`
}

type AccountExistsResponseDTO struct {
	Exists bool   `json:"exists"`
	Error  string `json:"err,omitempty"`
}
