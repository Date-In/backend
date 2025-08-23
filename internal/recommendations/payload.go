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
	ID                uint            `json:"id"`
	Name              string          `json:"name"`
	Age               uint            `json:"age"`
	Bio               *string         `json:"bio"`
	City              *string         `json:"city"`
	Children          *bool           `json:"children"`
	Height            *uint           `json:"height"`
	Sex               *ReferenceDto   `json:"sex"`
	ZodiacSign        *ReferenceDto   `json:"zodiac_sign"`
	Worldview         *ReferenceDto   `json:"worldview"`
	TypeOfDating      *ReferenceDto   `json:"type_of_dating"`
	Education         *ReferenceDto   `json:"education"`
	AttitudeToAlcohol *ReferenceDto   `json:"attitude_to_alcohol"`
	AttitudeToSmoking *ReferenceDto   `json:"attitude_to_smoking"`
	Status            *ReferenceDto   `json:"action"`
	Interests         []*ReferenceDto `json:"interests"`
	Avatar            *PhotoDto       `json:"avatar"`
	Gallery           []PhotoDto      `json:"gallery"`
}

type ReferenceDto struct {
	ID   *uint   `json:"id"`
	Name *string `json:"name"`
}

type PhotoDto struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}
