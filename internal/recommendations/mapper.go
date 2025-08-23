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
		ID:       user.ID,
		Name:     user.Name,
		Age:      user.Age,
		Bio:      user.Bio,
		City:     user.City,
		Children: user.Children,
		Height:   user.Height,
	}
	dto.Sex = &ReferenceDto{ID: &user.Sex.ID, Name: &user.Sex.Name}
	dto.ZodiacSign = &ReferenceDto{ID: &user.ZodiacSign.ID, Name: &user.ZodiacSign.Name}
	dto.Worldview = &ReferenceDto{ID: &user.Worldview.ID, Name: &user.Worldview.Name}
	dto.TypeOfDating = &ReferenceDto{ID: &user.TypeOfDating.ID, Name: &user.TypeOfDating.Name}
	dto.Education = &ReferenceDto{ID: &user.Education.ID, Name: &user.Education.Name}
	dto.AttitudeToAlcohol = &ReferenceDto{ID: &user.AttitudeToAlcohol.ID, Name: &user.AttitudeToAlcohol.Name}
	dto.AttitudeToSmoking = &ReferenceDto{ID: &user.AttitudeToSmoking.ID, Name: &user.AttitudeToSmoking.Name}
	dto.Status = &ReferenceDto{ID: &user.Status.ID, Name: &user.Status.Name}

	if user.Interests != nil && len(user.Interests) > 0 {
		dto.Interests = make([]*ReferenceDto, 0, len(user.Interests))
		for _, interest := range user.Interests {
			if interest != nil {
				dto.Interests = append(dto.Interests, &ReferenceDto{
					ID:   &interest.ID,
					Name: &interest.Name,
				})
			}
		}
	} else {
		dto.Interests = make([]*ReferenceDto, 0)
	}
	return dto
}
