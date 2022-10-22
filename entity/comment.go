package entity

import (
	"MyGram/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string    `gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	// validasi message
	if helper.IsEmpty(c.Message) {
		err = errors.New("Message can't be empty")
	}
	return
}
