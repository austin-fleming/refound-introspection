package domain

import (
	"fmt"
	shared "refound/internal/shared/domain"
)

// Aggregate of an account and its followers
type Follower struct {
	followerId    shared.IdVO
	followedSince shared.TimestampVO
	username      shared.UsernameVO
	avatarUrl     shared.UrlVO
}

type RawFollower struct {
	FollowerId    string
	FollowedSince string
	Username      string
	AvatarUrl     string
}

func ParseFollower(id string, followedSince string, username string, avatarUrl string) (Follower, error) {
	var follower Follower
	var validationErrors []error

	if parsedId, err := shared.ParseId(true, id); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		follower.followerId = parsedId
	}

	if parsedFollowedSince, err := shared.ParseTimestamp(true, followedSince); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		follower.followedSince = parsedFollowedSince
	}

	if parsedUsername, err := shared.ParseUsername(true, username); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		follower.username = parsedUsername
	}

	if parsedAvatarUrl, err := shared.ParseUrl(true, avatarUrl); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		follower.avatarUrl = parsedAvatarUrl
	}

	if len(validationErrors) > 0 {
		return Follower{}, fmt.Errorf("account had the following errors: %v", validationErrors)
	}

	return follower, nil
}

func (f Follower) ToValue() RawFollower {
	return RawFollower{
		FollowerId:    f.followerId.ToValue(),
		FollowedSince: f.followedSince.ToValue(),
		Username:      f.username.ToValue(),
		AvatarUrl:     f.avatarUrl.ToValue(),
	}
}
