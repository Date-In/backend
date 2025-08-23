package like

type LikeDto struct {
	ID   uint `json:"id"`
	User LikeUserDto
}

type LikeUserDto struct {
	ID      uint       `json:"id"`
	Name    string     `json:"name"`
	Age     uint       `json:"age"`
	Bio     *string    `json:"bio"`
	City    *string    `json:"city"`
	Avatar  *PhotoDto  `json:"avatar"`
	Gallery []PhotoDto `json:"gallery"`
}

type PhotoDto struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}
