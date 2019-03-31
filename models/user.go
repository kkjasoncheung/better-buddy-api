package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	PasswordDigest  string
	Email           string `json:"email"`
	BirthDay        string `json:"birthday"`
	Gender          string `json:"gender"`
	DisplayPhotoURL string `json:"display_photo_url"`
}
