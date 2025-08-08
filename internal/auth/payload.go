package auth

type RegisterRequestDto struct {
	Phone    string `validate:"required,min=8" json:"phone"`
	Name     string `validate:"required" json:"name"`
	Password string `validate:"required" json:"password"`
	Age      uint   `validate:"required,min=18" json:"age"`
	SexID    uint   `validate:"required" json:"sex_id"`
}

type LoginRequestDto struct {
	Phone    string `validate:"required" json:"phone"`
	Password string `validate:"required" json:"password"`
}

type RegisterResponseDto struct {
	ID    uint   `json:"id"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	SexID uint   `json:"sex_id"`
}
