package profile

import (
	"dating_service/internal/model"
)

func ToProfileResponseDto(user *model.User) *GetInfoResponseDto {
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
	dto := &GetInfoResponseDto{
		ID:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Age:      user.Age,
		City:     user.City,
		Bio:      user.Bio,
		Children: user.Children,
		Height:   user.Height,
		Avatar:   avatar,
		Gallery:  gallery,
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
