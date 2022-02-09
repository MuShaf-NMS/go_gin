package dto

type LoginDTO struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}
