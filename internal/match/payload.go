package match

type MatchDto struct {
	ID      uint `json:"id"`
	User1ID uint `json:"user1_id"`
	User2ID uint `json:"user2_id"`
}

type GetAllDto struct {
	Matches []MatchDto `json:"matches"`
}
