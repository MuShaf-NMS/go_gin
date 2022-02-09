package repository

import (
	"github.com/MuShaf-NMS/go_gin/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User)
	GetAllUsers() []models.User
	GetOneUser(uuid uuid.UUID) models.User
	CheckUser(username string) interface{}
	CheckUserEmailExist(username string, email string) bool
	UpdateUser(uuid uuid.UUID, user *models.User)
	DeleteUser(uuid uuid.UUID)
}

type userRepository struct {
	connection *gorm.DB
}

func (ur *userRepository) CreateUser(user *models.User) {
	ur.connection.Create(user)
}

func (ur *userRepository) GetAllUsers() []models.User {
	var users []models.User
	ur.connection.Find(&users)
	return users
}

func (ur *userRepository) GetOneUser(uuid uuid.UUID) models.User {
	var user models.User
	ur.connection.Where(models.User{UUID: uuid}).Take(&user)
	return user
}

func (ur *userRepository) CheckUser(username string) interface{} {
	var user models.User
	res := ur.connection.Where("username = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (ur *userRepository) CheckUserEmailExist(username string, email string) bool {
	var user models.User
	res := ur.connection.Where("username = ? or email = ?", username, email).Take(&user)
	return res.Error == nil
}

func (ur *userRepository) UpdateUser(uuid uuid.UUID, user *models.User) {
	ur.connection.Model(models.User{}).Where("uuid = ?", uuid).Updates(user)
}

func (ur *userRepository) DeleteUser(uuid uuid.UUID) {
	ur.connection.Delete(models.User{UUID: uuid})
}

func NewUserRepository(dbConnection *gorm.DB) UserRepository {
	return &userRepository{
		connection: dbConnection,
	}
}
