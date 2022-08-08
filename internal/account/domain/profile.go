package domain

import (
	"fmt"
	shared "refound/internal/shared/domain"
)

type Profile struct {
	username        shared.UsernameVO
	avatarUrl       shared.UrlVO
	isPrivate       bool
	followerCount   uint
	contractAddress shared.EthAddressVO  // public only
	joinedAt        shared.TimestampVO   // public only
	socialLinks     AccountSocialLinksVO // public only
	bio             AccountBioVO         // public only
}

type RawProfile struct {
	Username        string
	AvatarUrl       string
	IsPrivate       bool
	FollowerCount   uint
	ContractAddress string
	JoinedAt        string
	SocialLinks     []string
	Bio             string
}

func ParseProfile(raw RawProfile) (Profile, error) {
	var profile Profile
	var validationErrors []error

	if username, err := shared.ParseUsername(true, raw.Username); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		profile.username = username
	}

	if avatarUrl, err := shared.ParseUrl(false, raw.AvatarUrl); err != nil {
		validationErrors = append(validationErrors, err)
	} else {
		profile.avatarUrl = avatarUrl
	}

	profile.isPrivate = raw.IsPrivate

	profile.followerCount = raw.FollowerCount

	// Ignore these fields if account is set to private
	if !profile.isPrivate {
		if contractAddress, err := shared.ParseEthAddress(true, raw.ContractAddress); err != nil {
			validationErrors = append(validationErrors, err)
		} else {
			profile.contractAddress = contractAddress
		}

		if joinedAt, err := shared.ParseTimestamp(true, raw.JoinedAt); err != nil {
			validationErrors = append(validationErrors, err)
		} else {
			profile.joinedAt = joinedAt
		}

		if socialLinks, err := ParseAccountSocialLinks(false, raw.SocialLinks); err != nil {
			validationErrors = append(validationErrors, err)
		} else {
			profile.socialLinks = socialLinks
		}

		if bio, err := ParseAccountBio(false, raw.Bio); err != nil {
			validationErrors = append(validationErrors, err)
		} else {
			profile.bio = bio
		}
	}

	if len(validationErrors) != 0 {
		return Profile{}, fmt.Errorf("account had the following errors: %v", validationErrors)
	}

	return profile, nil
}

func (p Profile) ToValue() RawProfile {
	return RawProfile{
		Username:      p.username.ToValue(),
		AvatarUrl:     p.avatarUrl.ToValue(),
		IsPrivate:     p.isPrivate,
		FollowerCount: p.followerCount,
		// public only fields
		ContractAddress: p.contractAddress.ToValue(),
		JoinedAt:        p.joinedAt.ToValue(),
		SocialLinks:     p.socialLinks.ToValue(),
		Bio:             p.bio.ToValue(),
	}
}
