package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kkjasoncheung/better-buddy-api/db"
	"golang.org/x/crypto/bcrypt"
)

// User struct has one Companion.
type User struct {
	gorm.Model
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Username        string    `json:"username"`
	PasswordDigest  string    `json:"password_digest"`
	Email           string    `json:"email"`
	BirthDay        string    `json:"birthday"`
	Gender          string    `json:"gender"`
	DisplayPhotoURL string    `json:"display_photo_url"`
	CompanionID     uint      `json:"companionID"`
	Companion       Companion `json:"Companion"`
}

// CreateUser creates a new user and stores it in the database.
func (u User) CreateUser(firstName, lastName, username, password, email, birthday, gender, displayPhotoURL string) User {
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

	return user
}

// GetAllUsers returns all users.
func (u User) GetAllUsers() []User {
	users := []User{}
	db := db.GetDb()
	db.Find(&users)
	return users
}

// FindByID finds a user by id.
func (u User) FindByID(id int) *User {
	db := db.GetDb()
	user := new(User)
	db.First(&user, id)
	return user
}

// UpdateByID updates a user's attributes by id.
func (u User) UpdateByID(id int, fields map[string]string) *User {
	db := db.GetDb()

	user := new(User)
	user = user.FindByID(id)

	for key, value := range fields {
		switch key {
		case "FirstName":
			user.FirstName = value
		case "LastName":
			user.LastName = value
		case "Username":
			user.Username = value
		case "Password":
			newPassword := HashPassword(value)
			user.PasswordDigest = newPassword
		case "Email":
			user.Email = value
		case "BirthDay":
			user.BirthDay = value
		case "Gender":
			user.Gender = value
		case "DisplayPhotoUrl":
			user.DisplayPhotoURL = value
		}
	}
	db.Save(&user)

	return user
}

// DeleteByID deletes a user by id.
func (u User) DeleteByID(id int) *User {
	db := db.GetDb()
	user := new(User)
	user = user.FindByID(id)
	db.Delete(&user)
	return user
}

// GetCompanion returns the companion associated with this user.
func (u User) GetCompanion() *Companion {
	db := db.GetDb()
	companion := new(Companion)
	db.First(&companion, u.CompanionID)
	return companion
}

// ChangeCompanion changes the companion for the user.
func (u *User) ChangeCompanion(newID uint) {
	db := db.GetDb()
	newCompanion := new(Companion)
	if err := db.First(&newCompanion, newID).Error; err != nil {
		log.Fatal(err)
	}
}

// HashPassword hashes the password passed in with bcrypt.
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}
