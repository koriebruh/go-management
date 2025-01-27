package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=5"`
	Email    string `json:"email" validate:"email,required,min=15"`
}
