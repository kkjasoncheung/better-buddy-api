package models

import (
	"log"
	"strconv"

	"github.com/kkjasoncheung/better-buddy-api/errors"

	"github.com/jinzhu/gorm"
	"github.com/kkjasoncheung/better-buddy-api/db"
	"golang.org/x/crypto/bcrypt"
)

// User struct has one Companion.
type User struct {
	gorm.Model
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	BirthDay        string `json:"birthday"`
	Gender          string `json:"gender"`
	DisplayPhotoURL string `json:"display_photo_url"`
	PasswordDigest  string
	Companion       Companion `json:"Companion"`
}

// CreateUser creates a new user and stores it in the database.
func (u User) CreateUser(fields map[string]string) (User, error) {
	db := db.GetDb()

	user := User{
		FirstName:       fields["first_name"],
		LastName:        fields["last_name"],
		Username:        fields["username"],
		PasswordDigest:  HashPassword(fields["password"]),
		Email:           fields["email"],
		BirthDay:        fields["birthday"],
		Gender:          fields["gender"],
		DisplayPhotoURL: fields["display_photo_url"],
	}
	if err := db.Create(&user).Error; err != nil {
		log.Println(err)
		return User{}, err
	}
	return user, nil
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
	// TODO: Throw error if user not found.
	return user
}

// UpdateByID updates a user's attributes by id.
func (u User) UpdateByID(id uint, fields map[string]string) (*User, error) {
	db := db.GetDb()
	user := u.FindByID(id)

	for key, value := range fields {
		switch key {
		case "first_name":
			user.FirstName = value
		case "last_name":
			user.LastName = value
		case "username":
			user.Username = value
		case "new_password":
			// Verify if current user has correct pwd
			if HashPassword(fields["old_password"]) != user.PasswordDigest {
				return nil, errors.NewInvalidPasswordError()
			}
			newPassword := HashPassword(value)
			user.PasswordDigest = newPassword
		case "email":
			user.Email = value
		case "birthday":
			user.BirthDay = value
		case "gender":
			user.Gender = value
		case "display_photo_url":
			user.DisplayPhotoURL = value
		case "companion_id":
			if newVal, err := strconv.ParseUint(value, 10, 32); err == nil {
				_, err := user.ChangeCompanion(uint(newVal))
				if err != nil {
					log.Println(err)
					return nil, err
				}
			} else {
				return nil, err
			}
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
		log.Println(err)
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
		log.Println(err)
		return nil, err
	}

	newCompanion.UserID = u.ID
	u.Companion = *newCompanion
	db.Save(&newCompanion)
	db.Save(&u)
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
