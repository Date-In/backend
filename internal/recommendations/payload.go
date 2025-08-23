package recommendations

type GetRecommendationsDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type GetRecommendationsRes struct {
	User              UserForRecommendations
	PercentageOfMatch int
}

type UserForRecommendations struct {
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
