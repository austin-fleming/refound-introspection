package pg

type AccountDSO struct {
	Id                      string   `json:"id"`
	WalletAddress           string   `json:"wallet_address"`
	ContractAddress         string   `json:"contract_address"`
	Username                string   `json:"username"`
	SocialLinks             []string `json:"social_links"`
	AvatarUrl               string   `json:"avatar_url"`
	Bio                     string   `json:"bio"`
	IsPrivate               bool     `json:"is_private"`
	Email                   string   `json:"email"`
	EmailIsVerified         bool     `json:"email_is_verified"`
	EmailMarketingIsAllowed bool     `json:"email_marketing_is_allowed"`
	PhoneNumber             string   `json:"phone_number"`
	PhoneIsVerified         bool     `json:"phone_is_verified"`
	CreatedAt               string   `json:"created_at"`
	UpdatedAt               string   `json:"updated_at"`
	Nonce                   string   `json:"nonce"`
}
