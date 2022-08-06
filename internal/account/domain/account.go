package domain

import (
	"fmt"
	shared "refound/internal/shared/domain"
)

type Account struct {
	id                      shared.IdVO
	walletAddress           shared.EthAddressVO
	contractAddress         shared.EthAddressVO
	username                shared.UsernameVO
	socialLinks             AccountSocialLinksVO
	avatarUrl               shared.UrlVO
	bio                     AccountBioVO
	isPrivate               bool
	email                   shared.EmailVO
	emailIsVerified         bool
	emailMarketingIsAllowed bool
	phoneNumber             shared.PhoneNumberVO
	phoneIsVerified         bool
	createdAt               shared.TimestampVO
	updatedAt               shared.TimestampVO
	nonce                   AccountNonceVO
}

type RawAccount struct {
	Id                      string
	WalletAddress           string
	ContractAddress         string
	Username                string
	SocialLinks             []string
	AvatarUrl               string
	Bio                     string
	IsPrivate               bool
	Email                   string
	EmailIsVerified         bool
	EmailMarketingIsAllowed bool
	PhoneNumber             string
	PhoneIsVerified         bool
	CreatedAt               string
	UpdatedAt               string
	Nonce                   string
}

func ParseAccount(raw RawAccount) (Account, error) {
	var account Account
	var validationErrors []error

	if id, err := shared.ParseId(true, raw.Id); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.id = id
	}

	if walletAddress, err := shared.ParseEthAddress(true, raw.WalletAddress); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.walletAddress = walletAddress
	}

	// optional
	if contractAddress, err := shared.ParseEthAddress(false, raw.ContractAddress); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.contractAddress = contractAddress
	}

	if username, err := shared.ParseUsername(true, raw.Username); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.username = username
	}

	// optional
	if socialLinks, err := ParseAccountSocialLinks(false, raw.SocialLinks); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.socialLinks = socialLinks
	}

	// optional
	if avatarUrl, err := shared.ParseUrl(false, raw.AvatarUrl); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.avatarUrl = avatarUrl
	}

	// optional
	if bio, err := ParseAccountBio(false, raw.Bio); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.bio = bio
	}

	account.isPrivate = raw.IsPrivate || false

	// optional
	if email, err := shared.ParseEmail(false, raw.Email); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.email = email
	}

	account.emailIsVerified = raw.EmailIsVerified || false
	account.emailMarketingIsAllowed = raw.EmailMarketingIsAllowed || false

	// optional
	if phoneNumber, err := shared.ParsePhoneNumber(false, raw.PhoneNumber); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.phoneNumber = phoneNumber
	}

	account.phoneIsVerified = raw.PhoneIsVerified || false

	if createdAt, err := shared.ParseTimestamp(true, raw.CreatedAt); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.createdAt = createdAt
	}

	if updatedAt, err := shared.ParseTimestamp(true, raw.UpdatedAt); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.updatedAt = updatedAt
	}

	if nonce, err := ParseAccountNonce(true, raw.Nonce); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.nonce = nonce
	}

	if len(validationErrors) != 0 {
		return Account{}, fmt.Errorf("account had the following errors: %v", validationErrors)
	}

	return account, nil
}

func NewAccount(raw RawAccount) (Account, error) {
	if raw.Id == "" {
		raw.Id = shared.GenerateId().ToValue()
	}

	if raw.CreatedAt == "" {
		raw.CreatedAt = shared.GenerateTimestamp().ToValue()
	}

	if raw.UpdatedAt == "" {
		raw.UpdatedAt = shared.GenerateTimestamp().ToValue()
	}

	if raw.Nonce == "" {
		raw.Nonce = GenerateAccountNonce().ToValue()
	}

	return ParseAccount(raw)
}

func (a Account) ToValue() RawAccount {
	return RawAccount{
		Id:                      a.id.ToValue(),
		WalletAddress:           a.walletAddress.ToValue(),
		ContractAddress:         a.contractAddress.ToValue(),
		Username:                a.username.ToValue(),
		SocialLinks:             a.socialLinks.ToValue(),
		AvatarUrl:               a.avatarUrl.ToValue(),
		Bio:                     a.bio.ToValue(),
		IsPrivate:               a.isPrivate,
		Email:                   a.email.ToValue(),
		EmailIsVerified:         a.emailIsVerified,
		EmailMarketingIsAllowed: a.emailMarketingIsAllowed,
		PhoneNumber:             a.phoneNumber.ToValue(),
		PhoneIsVerified:         a.phoneIsVerified,
		CreatedAt:               a.createdAt.ToValue(),
		UpdatedAt:               a.updatedAt.ToValue(),
		Nonce:                   a.nonce.ToValue(),
	}
}

func (a Account) GetId() string {
	return a.id.ToValue()
}
