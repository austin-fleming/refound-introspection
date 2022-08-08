package account_service

import (
	accountRepo "refound/internal/account/repo"
	shared "refound/internal/shared/domain"
)

type AccountService interface {
	Exists(id shared.IdVO) (bool, error)
	Get(id shared.IdVO) (accountRepo)
}

type service struct {
	account         accountRepo.AccountRepo
	accountRelation accountRepo.AccountRelationRepo
}

func NewService(accountRepo accountRepo.AccountRepo, accountRelationRepo accountRepo.AccountRelationRepo) service {
	return service{
		account:         accountRepo,
		accountRelation: accountRelationRepo,
	}
}

