package recommendations

import "dating_service/internal/profile"

func ScoredUserToGetRecommendationResponse(scoredUser []ScoredUser) []GetRecommendationsRes {
	var res []GetRecommendationsRes
	for _, u := range scoredUser {
		res = append(res, GetRecommendationsRes{
			User:              *profile.ToProfileResponseDto(&u.Users),
			PercentageOfMatch: int(u.Score),
		})
	}
	return res
}
