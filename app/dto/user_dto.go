package dto

type CreateUser struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUser struct {
	Username string `json:"username" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdatePasswordUser struct {
	Password string `json:"password" binding:"required,min=8"`
}
