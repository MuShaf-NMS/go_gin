package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	Id          uint      `gorm:"id;primaryKey;autoIncrement" json:"-"`
	UUID        uuid.UUID `gorm:"uuid;type:char(36);not null" json:"uuid"`
	Name        string    `gorm:"name;size:100" json:"name"`
	Category    string    `gorm:"category;size:100" json:"category"`
	Description string    `gorm:"description" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	UserID      uuid.UUID `gorm:"userID;not null" json:"userID"`
}

func (todo *Todo) BeforeCreate(scope *gorm.DB) error {
	var err error
	todo.UUID, err = uuid.NewRandom()
	if err != nil {
		return err
	}
	return err
}
