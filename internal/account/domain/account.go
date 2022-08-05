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

type rawAccount struct {
	id                      string
	walletAddress           string
	contractAddress         string
	username                string
	socialLinks             []string
	avatarUrl               string
	bio                     string
	isPrivate               bool
	email                   string
	emailIsVerified         bool
	emailMarketingIsAllowed bool
	phoneNumber             string
	phoneIsVerified         bool
	createdAt               string
	updatedAt               string
	nonce                   string
}

func ParseAccount(raw rawAccount) (Account, error) {
	var account Account
	var validationErrors []error

	if id, err := shared.ParseId(true, raw.id); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.id = id
	}

	if walletAddress, err := shared.ParseEthAddress(true, raw.walletAddress); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.walletAddress = walletAddress
	}

	// optional
	if contractAddress, err := shared.ParseEthAddress(false, raw.contractAddress); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.contractAddress = contractAddress
	}

	if username, err := shared.ParseUsername(true, raw.username); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.username = username
	}

	// optional
	if socialLinks, err := ParseAccountSocialLinks(false, raw.socialLinks); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.socialLinks = socialLinks
	}

	// optional
	if avatarUrl, err := shared.ParseUrl(false, raw.avatarUrl); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.avatarUrl = avatarUrl
	}

	// optional
	if bio, err := ParseAccountBio(false, raw.bio); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.bio = bio
	}

	account.isPrivate = raw.isPrivate || false

	// optional
	if email, err := shared.ParseEmail(false, raw.email); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.email = email
	}

	account.emailIsVerified = raw.emailIsVerified || false
	account.emailMarketingIsAllowed = raw.emailMarketingIsAllowed || false

	// optional
	if phoneNumber, err := shared.ParsePhoneNumber(false, raw.phoneNumber); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.phoneNumber = phoneNumber
	}

	account.phoneIsVerified = raw.phoneIsVerified || false

	if createdAt, err := shared.ParseTimestamp(true, raw.createdAt); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.createdAt = createdAt
	}

	if updatedAt, err := shared.ParseTimestamp(true, raw.updatedAt); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.updatedAt = updatedAt
	}

	if nonce, err := ParseAccountNonce(true, raw.nonce); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		account.nonce = nonce
	}

	if len(validationErrors) != 0 {
		return Account{}, fmt.Errorf("account had the following errors: %v", validationErrors)
	}

	return account, nil
}

func NewAccount(raw rawAccount) (Account, error) {
	if raw.id == "" {
		raw.id = shared.GenerateId().ToValue()
	}

	if raw.createdAt == "" {
		raw.createdAt = shared.GenerateTimestamp().ToValue()
	}

	if raw.updatedAt == "" {
		raw.updatedAt = shared.GenerateTimestamp().ToValue()
	}

	if raw.nonce == "" {
		raw.nonce = GenerateAccountNonce().ToValue()
	}

	return ParseAccount(raw)
}

func (a Account) ToValue() rawAccount {
	return rawAccount{
		id:                      a.id.ToValue(),
		walletAddress:           a.walletAddress.ToValue(),
		contractAddress:         a.contractAddress.ToValue(),
		username:                a.username.ToValue(),
		socialLinks:             a.socialLinks.ToValue(),
		avatarUrl:               a.avatarUrl.ToValue(),
		bio:                     a.bio.ToValue(),
		isPrivate:               a.isPrivate,
		email:                   a.email.ToValue(),
		emailIsVerified:         a.emailIsVerified,
		emailMarketingIsAllowed: a.emailMarketingIsAllowed,
		phoneNumber:             a.phoneNumber.ToValue(),
		phoneIsVerified:         a.phoneIsVerified,
		createdAt:               a.createdAt.ToValue(),
		updatedAt:               a.updatedAt.ToValue(),
		nonce:                   a.nonce.ToValue(),
	}
}
