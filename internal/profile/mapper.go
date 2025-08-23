package profile

import (
	"dating_service/internal/model"
)

func ToUserInfoResponseDto(user *model.User) *GetInfoResponseDto {
	if user == nil {
		return nil
	}

	gallery := mapPhotos(user.Photos)
	interests := mapInterests(user.Interests)

	return &GetInfoResponseDto{
		ID:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Age:      user.Age,
		City:     user.City,
		Bio:      user.Bio,
		Children: user.Children,
		Height:   user.Height,

		Avatar: &PhotoDto{
			ID:  user.Avatar.ID,
			Url: user.Avatar.Url,
		},
		Gallery:   gallery,
		Interests: interests,

		Sex:               mapReference(&user.Sex),
		ZodiacSign:        mapReference(&user.ZodiacSign),
		Worldview:         mapReference(&user.Worldview),
		TypeOfDating:      mapReference(&user.TypeOfDating),
		Education:         mapReference(&user.Education),
		AttitudeToAlcohol: mapReference(&user.AttitudeToAlcohol),
		AttitudeToSmoking: mapReference(&user.AttitudeToSmoking),
		Status:            mapReference(&user.Status),
	}
}

func ToUserProfileResponseDto(user *model.User) *GetProfileResponseDto {
	if user == nil {
		return nil
	}

	gallery := mapPhotos(user.Photos)
	interests := mapInterests(user.Interests)

	return &GetProfileResponseDto{
		ID:       user.ID,
		Name:     user.Name,
		Age:      user.Age,
		City:     user.City,
		Bio:      user.Bio,
		Children: user.Children,
		Height:   user.Height,

		Avatar: &PhotoDto{
			ID:  user.Avatar.ID,
			Url: user.Avatar.Url,
		},
		Gallery:   gallery,
		Interests: interests,

		Sex:               mapReference(&user.Sex),
		ZodiacSign:        mapReference(&user.ZodiacSign),
		Worldview:         mapReference(&user.Worldview),
		TypeOfDating:      mapReference(&user.TypeOfDating),
		Education:         mapReference(&user.Education),
		AttitudeToAlcohol: mapReference(&user.AttitudeToAlcohol),
		AttitudeToSmoking: mapReference(&user.AttitudeToSmoking),
		Status:            mapReference(&user.Status),
	}
}

func mapPhotos(photos []*model.Photo) []PhotoDto {
	gallery := make([]PhotoDto, 0)

	if photos == nil {
		return gallery
	}
	for _, p := range photos {
		gallery = append(gallery, PhotoDto{
			ID:  p.ID,
			Url: p.Url,
		})
	}
	return gallery
}

type Referable interface {
	GetID() uint
	GetName() string
}

func mapReference(ref Referable) *ReferenceDto {
	if ref == nil {
		return &ReferenceDto{}
	}
	id := ref.GetID()
	name := ref.GetName()
	return &ReferenceDto{
		ID:   &id,
		Name: &name,
	}
}

func mapInterests(interests []*model.Interest) []*ReferenceDto {
	if interests == nil || len(interests) == 0 {
		return make([]*ReferenceDto, 0)
	}

	dtoList := make([]*ReferenceDto, 0, len(interests))

	for _, interest := range interests {
		if interest != nil {
			dtoList = append(dtoList, mapReference(interest))
		}
	}
	return dtoList
}
