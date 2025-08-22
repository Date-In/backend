package profile

import "dating_service/internal/photo"

type GetInfoResponseDto struct {
	ID                uint             `json:"id"`
	Name              string           `json:"name"`
	Phone             string           `json:"phone"`
	Age               uint             `json:"age"`
	Bio               *string          `json:"bio"`
	City              *string          `json:"city"`
	Children          *bool            `json:"children"`
	Height            *uint            `json:"height"`
	Sex               *ReferenceDto    `json:"sex"`
	ZodiacSign        *ReferenceDto    `json:"zodiac_sign"`
	Worldview         *ReferenceDto    `json:"worldview"`
	TypeOfDating      *ReferenceDto    `json:"type_of_dating"`
	Education         *ReferenceDto    `json:"education"`
	AttitudeToAlcohol *ReferenceDto    `json:"attitude_to_alcohol"`
	AttitudeToSmoking *ReferenceDto    `json:"attitude_to_smoking"`
	Status            *ReferenceDto    `json:"action"`
	Interests         []*ReferenceDto  `json:"interests"`
	Avatar            *photo.PhotoDto  `json:"avatar"`
	Gallery           []photo.PhotoDto `json:"gallery"`
}

type ReferenceDto struct {
	ID   *uint   `json:"id"`
	Name *string `json:"name"`
}

type UpdateInfoRequestDto struct {
	Name                *string `json:"name"`
	Age                 *uint   `json:"age"`
	Bio                 *string `json:"bio"`
	City                *string `json:"city"`
	Children            *bool   `json:"children"`
	Height              *uint   `json:"height"`
	SexId               *uint   `json:"sex_id"`
	ZodiacSignId        *uint   `json:"zodiac_sign_id"`
	WorldviewId         *uint   `json:"worldview_id"`
	TypeOfDatingId      *uint   `json:"type_of_dating_id"`
	EducationId         *uint   `json:"education_id"`
	AttitudeToAlcoholId *uint   `json:"attitude_to_alcohol_id"`
	AttitudeToSmokingId *uint   `json:"attitude_to_smoking_id"`
}

type UpdateInterestRequestDto struct {
	InterestIDs []uint `json:"interests" validate:"required"`
}
