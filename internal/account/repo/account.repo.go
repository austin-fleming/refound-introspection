package repo

import (
	accountDomain "refound/internal/account/domain"
	shared "refound/internal/shared/domain"
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
	Exists(id shared.IdVO) (bool, error)
	Get(id shared.IdVO) (accountDomain.Account, error)
	ExistsByWalletAddress(walletAddress shared.EthAddressVO) (bool, error)
	GetByWalletAddress(walletAddress shared.EthAddressVO) (accountDomain.Account, error)
	GetNonceByWalletAddress(WalletAddress shared.EthAddressVO) (accountDomain.AccountNonceVO, error)
	Replace(account accountDomain.Account) (accountDomain.Account, error)
	UpdateBeneficiary() (error)
	Delete(account accountDomain.Account) (accountDomain.Account, error)
}
