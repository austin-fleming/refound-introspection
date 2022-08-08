package pg

type InitiativeMemberRelationDSO struct {
	Initiative     string `json:"initiative"`
	Follower       string `json:"follower"`
	MembershipRole string `json:"membership_role"`
	CreatedAt      string `json:"created_at"`
}
