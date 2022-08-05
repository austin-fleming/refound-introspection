package repo

import (
	pg "refound/db/pg"
)

const (
	ErrAccountRelGetFollowers = "account relation: could not get followers"
	ErrAccountRelGetFollowing = "account relation: could not get following"
	ErrAccountRelCreate       = "account relation: could not create"
	ErrAccountRelDelete       = "account relation: could not delete"
)

type AccountRelationRepo interface {
	GetFollowers(accountId string) ([]pg.AccountFollowRelationDSO, error)
	GetFollowing(accountId string) ([]pg.AccountFollowRelationDSO, error)
	Create(accountId string, followeeId string) (pg.AccountFollowRelationDSO, error)
	Delete(accountId string, followeeId string) (pg.AccountFollowRelationDSO, error)
}

func MakeAccountRelationRepo() AccountRelationRepo {
	return AccountRelationRepo{}
}
