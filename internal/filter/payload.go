package filter

type CreateFilterDto struct {
	MinAge   uint   `json:"min_age"`
	MaxAge   uint   `json:"max_age"`
	SexID    uint   `json:"sex_id"`
	Location string `json:"location"`
}

type UpdateFilterDto struct {
	MinAge   *uint   `json:"min_age"`
	MaxAge   *uint   `json:"max_age"`
	SexId    *uint   `json:"sex_id"`
	Location *string `json:"location"`
}

type GetFilterDto struct {
	MinAge   uint   `json:"min_age"`
	MaxAge   uint   `json:"max_age"`
	SexID    uint   `json:"sex_id"`
	Location string `json:"location"`
}
