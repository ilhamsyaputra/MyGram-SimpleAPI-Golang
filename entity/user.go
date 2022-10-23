package entity

import (
	"MyGram/helper"
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string        `gorm:"primaryKey" json:"id"`
	Username     string        `json:"username"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Age          int           `json:"age"`
	CreatedAt    time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time     `gorm:"column:updated_at" json:"updated_at"`
	SocialMedias []SocialMedia `gorm:"foreignKey:ID"`
	Comments     []Comment     `gorm:"foreignKey:ID"`
	Photos       []Photo       `gorm:"foreignKey:ID"`
}

type UserRegisterResponse struct {
	ID       string
	Username string
	Email    string
	Age      int
}

type UserUpdateResponse struct {
	ID        string
	Username  string
	Email     string
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	// validasi field email
	if helper.IsEmpty(u.Email) {
		err = errors.New("Email can't be empty")
	}

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(u.Email) {
		err = errors.New("Email not valid")
	}

	// Validasi username
	if helper.IsEmpty(u.Username) {
		err = errors.New("Username can't be empty")
	}

	// Validasi password
	if helper.IsEmpty(u.Password) {
		err = errors.New("Password can't be empty")
	} else if len(u.Password) < 6 {
		err = errors.New("Password must have at least 6 characters length")
	}

	// validasi age
	if helper.IsEmpty(u.Age) {
		err = errors.New("Age can't be empty")
	}

	if u.Age < 8 {
		err = errors.New("Age must be at least 8 years old")
	}

	u.ID = uuid.New().String()
	u.Password = helper.HashPass(u.Password)

	return
}
