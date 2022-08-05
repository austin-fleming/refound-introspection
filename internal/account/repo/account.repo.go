package repo

import (
	accountDomain "refound/internal/account/domain"
)

const (
	ErrAccountRepoExists                  = "account: could not check existence"
	ErrAccountRepoGet                     = "account: could not get"
	ErrAccountRepoGetNonceByWalletAddress = "account: could not get nonce"
	ErrAccountRepoUpdate                  = "account: could not update"
	ErrAccountRepoUpdateBeneficiary       = "account: could not update beneficiary"
	ErrAccountDelete                      = "account: could not delete"
)

type AccountRepo interface {
	Exists(id string) (bool, error)
	Get(id string) (accountDomain.Account, error)
	ExistsByWalletAddress(walletAddress string) (bool, error)
	GetByWalletAddress(walletAddress string) (accountDomain.Account, error)
	GetNonceByWalletAddress(WalletAddress string) (string, error)
	Update(account accountDomain.Account) (accountDomain.Account, error)
	UpdateBeneficiary(account accountDomain.Account) (bool, error)
	Delete(account accountDomain.Account) (accountDomain.Account, error)
}

func MakeAccountRepo() AccountRepo {
	return AccountRepo{}
}
