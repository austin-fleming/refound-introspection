package pg

type InitiativeVoteDSO struct {
	Initiative string `json:"initiative"`
	Voter      string `json:"voter"`
	Vote       int8   `json:"vote"`
	CreatedAt  string `json:"created_at"`
}
