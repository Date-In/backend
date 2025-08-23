package recommendations

import (
	"dating_service/internal/model"
)

func ScoredUserToGetRecommendationResponse(scoredUser []ScoredUser) []GetRecommendationsRes {
	var res []GetRecommendationsRes
	for _, u := range scoredUser {
		res = append(res, GetRecommendationsRes{
			User:              *UserToUserForRecommendations(&u.User),
			PercentageOfMatch: int(u.Score),
		})
	}
	return res
}

func UserToUserForRecommendations(user *model.User) *UserForRecommendations {
	if user == nil {
		return nil
	}
	photoUser := user.Photos
	avatar := &PhotoDto{}
	var gallery []PhotoDto
	for _, p := range photoUser {
		if p.IsAvatar {
			avatar.Url = p.Url
			avatar.ID = p.ID
		} else {
			gallery = append(gallery, PhotoDto{
				ID:  p.ID,
				Url: p.Url,
			})
		}
	}
	dto := &UserForRecommendations{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
		Bio:  user.Bio,
		City: user.City,
	}
	return dto
}
