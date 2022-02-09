package service

import (
	"github.com/MuShaf-NMS/go_gin/app/dto"
	"github.com/MuShaf-NMS/go_gin/app/models"
	"github.com/MuShaf-NMS/go_gin/app/repository"
	"github.com/google/uuid"
)

type UserService interface {
	GetUser(uuid uuid.UUID) models.User
	UpdateUser(uuid uuid.UUID, user dto.UpdateUser)
	UpdatePassword(uuid uuid.UUID, user dto.UpdatePasswordUser)
}

type userService struct {
	repository repository.UserRepository
}

func (us *userService) GetUser(uuid uuid.UUID) models.User {
	user := us.repository.GetOneUser(uuid)
	return user
}

func (us *userService) UpdateUser(uuid uuid.UUID, user dto.UpdateUser) {
	userUpdate := models.User{
		Username: user.Username,
		Email:    user.Email,
	}
	us.repository.UpdateUser(uuid, &userUpdate)
}

func (us *userService) UpdatePassword(uuid uuid.UUID, user dto.UpdatePasswordUser) {
	passUpdate := models.User{
		Password: user.Password,
	}
	us.repository.UpdateUser(uuid, &passUpdate)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		repository: userRepo,
	}
}
