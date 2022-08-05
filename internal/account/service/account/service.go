package account_service

import (
	accountRepo "refound/internal/account/repo"
)

type AccountService interface {
	Exists()
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
