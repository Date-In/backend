package profile

type GetInfoResponseDto struct {
	ID                uint            `json:"id"`
	Name              string          `json:"name"`
	Phone             string          `json:"phone"`
	Age               uint            `json:"age"`
	Bio               *string         `json:"bio"`
	Children          *bool           `json:"children"`
	Height            *uint           `json:"height"`
	Sex               *ReferenceDto   `json:"sex"`
	ZodiacSign        *ReferenceDto   `json:"zodiac_sign"`
	Worldview         *ReferenceDto   `json:"worldview"`
	TypeOfDating      *ReferenceDto   `json:"type_of_dating"`
	Education         *ReferenceDto   `json:"education"`
	AttitudeToAlcohol *ReferenceDto   `json:"attitude_to_alcohol"`
	AttitudeToSmoking *ReferenceDto   `json:"attitude_to_smoking"`
	Status            *ReferenceDto   `json:"status"`
	Interests         []*ReferenceDto `json:"interests"`
}

type ReferenceDto struct {
	ID   *uint   `json:"id"`
	Name *string `json:"name"`
}
