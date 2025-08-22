package recommendations

import "dating_service/internal/profile"

type GetRecommendationsDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type GetRecommendationsRes struct {
	User              profile.GetInfoResponseDto
	PercentageOfMatch int
}
