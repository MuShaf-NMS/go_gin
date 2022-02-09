package service

import (
	"github.com/MuShaf-NMS/go_gin/app/dto"
	"github.com/MuShaf-NMS/go_gin/app/models"
	"github.com/MuShaf-NMS/go_gin/app/repository"
)

type AuthService interface {
	RegisterUser(user dto.CreateUser) models.User
	VerifyUser(user dto.LoginDTO) interface{}
	CheckUser(user dto.CreateUser) bool
}

type authService struct {
	repository repository.UserRepository
}

func (as *authService) RegisterUser(user dto.CreateUser) models.User {
	userCreate := models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	as.repository.CreateUser(&userCreate)
	return userCreate
}

func (as *authService) VerifyUser(user dto.LoginDTO) interface{} {
	res := as.repository.CheckUser(user.Username)
	if v, ok := res.(models.User); ok {
		if verifyPassword(user.Password, v.Password) {
			return res
		}
		return false
	}
	return false
}

func (as *authService) CheckUser(user dto.CreateUser) bool {
	res := as.repository.CheckUserEmailExist(user.Username, user.Email)
	return res
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		repository: userRepo,
	}
}

func verifyPassword(plain string, hashed string) bool {
	return plain == hashed
}
