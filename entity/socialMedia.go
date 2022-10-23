package entity

import (
	"MyGram/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	// validasi name
	if helper.IsEmpty(s.Name) {
		err = errors.New("Name can't be empty")
	}

	// validasi url
	if helper.IsEmpty(s.SocialMediaUrl) {
		err = errors.New("url can't be empty")
	}

	return
}
