package pg

type InitiativeDSO struct {
	Id              string `json:"id"`
	ContractAddress string `json:"contract_address"`
	Title           string `json:"title"`
	Body            string `json:"body"`
	CoverImage      string `json:"cover_image"`
	TargetDate      string `json:"target_date"`
	TargetFunds     uint64 `json:"target_funds"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
