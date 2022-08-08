package implementation

import (
	"database/sql"
	"errors"
	"fmt"
	pg "refound/internal/shared/db/pg"
	account "refound/internal/account/domain"
	shared "refound/internal/shared/domain"
)

type accountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) accountRepo {
	return accountRepo{db}
}

// -----------------
// HELPERS
// -----------------

func handlePqError(err error) error {
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return errors.New("No rows returned")
	default:
		return fmt.Errorf("unknown error: %s", err.Error())
	}
}

func dsoToEntity(dso pg.AccountDSO) (account.Account, error) {
	return account.ParseAccount(account.RawAccount{
		Id:                      dso.Id,
		WalletAddress:           dso.WalletAddress,
		ContractAddress:         dso.ContractAddress,
		Username:                dso.Username,
		SocialLinks:             dso.SocialLinks,
		AvatarUrl:               dso.AvatarUrl,
		Bio:                     dso.Bio,
		IsPrivate:               dso.IsPrivate,
		Email:                   dso.Email,
		EmailIsVerified:         dso.EmailIsVerified,
		EmailMarketingIsAllowed: dso.EmailMarketingIsAllowed,
		PhoneNumber:             dso.PhoneNumber,
		PhoneIsVerified:         dso.PhoneIsVerified,
		CreatedAt:               dso.CreatedAt,
		UpdatedAt:               dso.UpdatedAt,
		Nonce:                   dso.Nonce,
	})
}

// -----------------
// METHODS
// -----------------

func (ar *accountRepo) Exists(id shared.IdVO) (bool, error) {
	var exists bool
	queryErr := ar.db.QueryRow(`SELECT exists (SELECT 1 FROM account WHERE id = $1)`, id.ToValue()).Scan(&exists)
	if queryErr != nil {
		// TODO: needs custom errors, may return pqError
		return false, handlePqError(queryErr)
	}

	return exists, nil
}

func (ar *accountRepo) Get(id shared.IdVO) (account.Account, error) {
	var dso pg.AccountDSO

	queryErr := ar.db.QueryRow(`SELECT * FROM account WHERE id = $1`, id.ToValue()).Scan(&dso)
	if queryErr != nil {
		return account.Account{}, handlePqError(queryErr)
	}

	entity, parseErr := account.ParseAccount(account.RawAccount{
		Id:                      dso.Id,
		WalletAddress:           dso.WalletAddress,
		ContractAddress:         dso.ContractAddress,
		Username:                dso.Username,
		SocialLinks:             dso.SocialLinks,
		AvatarUrl:               dso.AvatarUrl,
		Bio:                     dso.Bio,
		IsPrivate:               dso.IsPrivate,
		Email:                   dso.Email,
		EmailIsVerified:         dso.EmailIsVerified,
		EmailMarketingIsAllowed: dso.EmailMarketingIsAllowed,
		PhoneNumber:             dso.PhoneNumber,
		PhoneIsVerified:         dso.PhoneIsVerified,
		CreatedAt:               dso.CreatedAt,
		UpdatedAt:               dso.UpdatedAt,
		Nonce:                   dso.Nonce,
	})
	if parseErr != nil {
		return account.Account{}, nil
	}

	return entity, nil
}

func (ar *accountRepo) ExistsByWalletAddress(wallet shared.EthAddressVO) (bool, error) {
	var exists bool
	queryErr := ar.db.QueryRow(`SELECT exists (SELECT 1 FROM account WHERE wallet_address = $1)`, wallet.ToValue()).Scan(&exists)
	if queryErr != nil {
		// TODO: needs custom errors, may return pqError
		return false, handlePqError(queryErr)
	}

	return exists, nil
}

func (ar *accountRepo) GetByWalletAddress(wallet shared.EthAddressVO) (account.Account, error) {
	var dso pg.AccountDSO

	queryErr := ar.db.QueryRow(`SELECT * FROM account WHERE wallet_address = $1`, wallet.ToValue()).Scan(&dso)
	if queryErr != nil {
		return account.Account{}, handlePqError(queryErr)
	}

	entity, parseErr := dsoToEntity(dso)
	if parseErr != nil {
		return account.Account{}, nil
	}

	return entity, nil
}

func (ar *accountRepo) GetNonceByWalletAddress(walletAddress shared.EthAddressVO) (account.AccountNonceVO, error) {
	var rawNonce string

	queryErr := ar.db.QueryRow(`SELECT nonce FROM account WHERE wallet_address = $1`, walletAddress.ToValue()).Scan(&rawNonce)
	if queryErr != nil {
		return account.AccountNonceVO{}, handlePqError(queryErr)
	}

	nonce, parseErr := account.ParseAccountNonce(true, rawNonce)
	if parseErr != nil {
		return account.AccountNonceVO{}, parseErr
	}

	return nonce, nil
}

func (ar *accountRepo) Replace(acc account.Account) (account.Account, error) {
	var updatedDSO pg.AccountDSO

	query := `
		UPDATE
			account 
		SET 
			id=$2, 
			wallet_address=$3, 
			contract_address=$4, 
			username=$5, 
			social_links=$6, 
			avatar_url=$7,
			bio=$8,
			is_private=$9,
			email=$10,
			email_is_verified=$11,
			email_marketing_is_allowed=$12,
			phone_number=$13,
			phone_is_verified=$14,
			created_at=$15,
			updated_at=$16,
			nonce=$17
		WHERE
			id = $1
	`

	v := acc.ToValue()

	queryErr := ar.db.QueryRow(
		query,
		v.Id,
		v.WalletAddress,
		v.ContractAddress,
		v.Username,
		v.SocialLinks,
		v.AvatarUrl,
		v.Bio,
		v.IsPrivate,
		v.Email,
		v.EmailIsVerified,
		v.EmailMarketingIsAllowed,
		v.PhoneNumber,
		v.PhoneIsVerified,
		v.CreatedAt,
		v.UpdatedAt,
		v.Nonce,
	).Scan(&updatedDSO)
	if queryErr != nil {
		return account.Account{}, handlePqError(queryErr)
	}

	entity, mapErr := dsoToEntity(updatedDSO)
	if mapErr != nil {
		return account.Account{}, handlePqError(mapErr)
	}

	return entity, nil
}

// TODO
func (ar *accountRepo) UpdateBeneficiary() error {
	return errors.New("Not implemented")
}

func (ar *accountRepo) Delete(acc account.Account) (account.Account, error) {
	// TODO: use "isDeleted" instead of actual delete for now
	id := acc.GetId()

	var deleted pg.AccountDSO
	query := "DELETE FROM account WHERE id = $1"
	queryErr := ar.db.QueryRow(query, id).Scan(&deleted)
	if queryErr != nil {
		return account.Account{}, handlePqError(queryErr)
	}

	// TODO: in theory, delete could succeed, but this fails creating a disparity in the program.
	// Rework interface or do a more thorough check.
	deletedEntity, mapErr := dsoToEntity(deleted)
	if mapErr != nil {
		return account.Account{}, mapErr
	}

	return deletedEntity, nil
}
