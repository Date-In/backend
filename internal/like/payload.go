package like

type LikeDto struct {
	ID       uint `json:"id"`
	UserID   uint `json:"user_id"`
	TargetID uint `json:"target_id"`
}

type GetLikeDto struct {
	Likes []LikeDto `json:"likes"`
}
