package implementation

import (
	"database/sql"
	"refound/internal/account/domain"
	persistence_utils "refound/internal/account/utils/persistence"
	shared "refound/internal/shared/domain"
)

type followerRepo struct {
	db *sql.DB
}

func NewFollowerRepo(db *sql.DB) followerRepo {
	return followerRepo{db}
}

// -----------------
// HELPERS
// -----------------

func (fr *followerRepo) GetFollowers(id shared.IdVO, limit, offset uint) ([]domain.Follower, error) {
	var followers []domain.Follower

	type JoinedFollower struct {
		Follower                string   `json:"follower"`
		Followee                string   `json:"followee"`
		CreatedAt               string   `json:"created_at"`
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
		UpdatedAt               string   `json:"updated_at"`
		Nonce                   string   `json:"nonce"`
	}

	query := `
			SELECT * 
			FROM account_follow_relation 
			WHERE followee = $1 
			ORDER BY created_at ASC 
			OFFSET $3 LIMIT $2
			INNER JOIN account
			ON account_follow_relation.follower_id = account.id
		`

	rows, queryErr := fr.db.Query(query, id.ToValue(), limit, offset)
	if queryErr != nil {
		return []domain.Follower{}, persistence_utils.HandlePqError(queryErr)
	}

	defer rows.Close()
	for rows.Next() {
		var dso JoinedFollower

		rowErr := rows.Scan(&dso)
		if rowErr != nil {
			return []domain.Follower{}, persistence_utils.HandlePqError(rowErr)
		}

		parsedFollower, parseErr := domain.ParseFollower(dso.Id, dso.CreatedAt, dso.Username, dso.AvatarUrl)
		if parseErr != nil {
			return []domain.Follower{}, parseErr
		}

		followers = append(followers, parsedFollower)
	}
	if rows.Err() != nil {
		return []domain.Follower{}, rows.Err()
	}

	return followers, nil
}
