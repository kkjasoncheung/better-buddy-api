package models

import (
	"log"

	"github.com/kkjasoncheung/better-buddy-api/errors"

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
func (u User) FindByID(id uint) *User {
	db := db.GetDb()
	user := new(User)
	db.First(&user, id)
	return user
}

// UpdateByID updates a user's attributes by id.
func (u User) UpdateByID(id uint, fields map[string]string) (*User, error) {
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
			// Verify if current user has correct pwd
			if HashPassword(fields["OldPassword"]) != user.PasswordDigest {
				return nil, errors.NewInvalidPasswordError()
			}
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

	return user, nil
}

// DeleteByID deletes a user by id.
func (u User) DeleteByID(id uint) *User {
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
	db.Model(&u).Related(&companion, "Companion")
	return companion
}

// ChangeCompanionByID changes the companion for the user based on userID.
func (u User) ChangeCompanionByID(userID uint, newID uint) (*User, error) {
	db := db.GetDb()

	user := new(User)
	user = user.FindByID(userID)

	newCompanion := new(Companion)
	if err := db.First(&newCompanion, newID).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}

	newCompanion.UserID = userID
	user.Companion = *newCompanion
	db.Save(&newCompanion)
	db.Save(&user)
	return user, nil
}

// ChangeCompanion changes the companion for the user.
func (u *User) ChangeCompanion(newID uint) (*Companion, error) {
	db := db.GetDb()
	newCompanion := new(Companion)
	if err := db.First(&newCompanion, newID).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}

	newCompanion.UserID = u.ID
	db.Save(&newCompanion)
	return newCompanion, nil
}

// HashPassword hashes the password passed in with bcrypt.
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}
