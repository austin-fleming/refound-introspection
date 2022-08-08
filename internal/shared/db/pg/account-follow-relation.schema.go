package pg

type AccountFollowRelationDSO struct {
	Follower  string `json:"follower"`
	Followee  string `json:"followee"`
	CreatedAt string `json:"created_at"`
}
