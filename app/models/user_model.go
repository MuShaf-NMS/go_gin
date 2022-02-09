package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type User struct {
	Id       uint      `gorm:"id;primaryKey;autoIncrement" json:"-"`
	UUID     uuid.UUID `gorm:"uuid;type:char(36);not null" json:"uuid"`
	Username string    `gorm:"username;size:100;not null" json:"username"`
	Password string    `gorm:"password;size:100;not null" json:"-"`
	Email    string    `gorm:"email;size:100;not null" json:"email"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	var err error
	user.UUID, err = uuid.NewRandom()
	if err != nil {
		return err
	}
	return err
}
