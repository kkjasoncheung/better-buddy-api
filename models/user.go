package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kkjasoncheung/better-buddy-api/db"
	"golang.org/x/crypto/bcrypt"
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

// GetAllUsers returns all users.
func (u User) GetAllUsers() []User {
	users := []User{}
	// Query the DB
	// db := db.GetDb()
	return users
}

// CreateUser creates a new user and stores it in the database.
func (u User) CreateUser(firstName, lastName, username, password, email, birthday, gender, displayPhotoURL string) {
	db := db.GetDb()

	user := User{
		FirstName:       firstName,
		LastName:        lastName,
		Username:        username,
		PasswordDigest:  HashPassword(password),
		Email:           email,
		BirthDay:        birthday,
		Gender:          gender,
		DisplayPhotoURL: displayPhotoURL,
	}
	db.Create(&user)
}

// HashPassword hashes the password passed in with bcrypt.
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}
