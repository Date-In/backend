package recommendations

type GetRecommendationsDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
