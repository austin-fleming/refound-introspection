package account_service

func (service service) Exists(request AccountExistsRequestDTO) (AccountExistsResponseDTO, error) {
	result, err := service.account.Exists(request.Id)
	if err != nil {
		return AccountExistsResponseDTO{}, err
	}

	return AccountExistsResponseDTO{Exists: result}, nil
}
