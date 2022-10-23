package entity

import (
	"MyGram/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `gorm:"foreignKey:PhotoID"`
}

type PostPhotoResponse struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	// validasi title
	if helper.IsEmpty(p.Title) {
		err = errors.New("Title can't be empty")
	}

	if helper.IsEmpty(p.PhotoURL) {
		err = errors.New("photo url can't be empty")
	}

	return
}
