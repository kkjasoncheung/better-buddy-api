package models

import (
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kkjasoncheung/better-buddy-api/db"
)

// Companion struct belongs to a User.
type Companion struct {
	gorm.Model
	Name               string  `json:"name"`
	Hunger             float32 `json:"hunger"`
	Tiredness          float32 `json:"tiredness"`
	Happiness          float32 `json:"happiness"`
	IntellectPotential float32 `json:"intellect_potential"`
	UserID             uint    `json:"UserID"`
}

// CreateCompanion creates a companion.
func (c Companion) CreateCompanion(name string, hunger, tiredness, happiness, intellectualPotential float32, userID uint) Companion {
	db := db.GetDb()

	companion := Companion{
		Name:               name,
		Hunger:             hunger,
		Tiredness:          tiredness,
		Happiness:          happiness,
		IntellectPotential: intellectualPotential,
		UserID:             userID,
	}
	db.Create(&companion)

	return companion
}

// GetAllCompanions gets all companions in database.
func (c Companion) GetAllCompanions() []Companion {
	db := db.GetDb()
	companions := []Companion{}
	db.Find(&companions)
	return companions
}

// FindByID finds a companion by ID.
func (c Companion) FindByID(id uint) *Companion {
	db := db.GetDb()
	companion := new(Companion)
	db.First(&companion, id)
	return companion
}

// UpdateByID updates the companion by id with the given key value pairs in fields.
func (c Companion) UpdateByID(id uint, fields map[string]string) (*Companion, error) {
	db := db.GetDb()

	companion := new(Companion)
	companion = companion.FindByID(id)

	for key, value := range fields {
		switch key {
		case "Name":
			companion.Name = value
		case "Hunger":
			if f, err := strconv.ParseFloat(value, 32); err == nil {
				companion.Hunger = float32(f)
			} else {
				return nil, err
			}

		case "Tiredness":
			if f, err := strconv.ParseFloat(value, 32); err == nil {
				companion.Tiredness = float32(f)
			} else {
				return nil, err
			}
		case "Happiness":
			if f, err := strconv.ParseFloat(value, 32); err == nil {
				companion.Happiness = float32(f)
			} else {
				return nil, err
			}
		case "IntellectPotential":
			if f, err := strconv.ParseFloat(value, 32); err == nil {
				companion.IntellectPotential = float32(f)
			} else {
				return nil, err
			}
		case "UserID":
			if f, err := strconv.ParseUint(value, 10, 32); err == nil {
				companion.UserID = uint(f)
			} else {
				return nil, err
			}
		}
	}

	return companion, nil
}

// GetUser returns the user associated with this companion.
func (c Companion) GetUser() *User {
	db := db.GetDb()
	user := new(User)
	db.First(&user, c.UserID)
	return user
}

// ChangeUserByID changes the user id for a companion given companionID.
func (c Companion) ChangeUserByID(companionID uint, newID uint) (*Companion, error) {
	db := db.GetDb()

	newUser := new(User)
	newUser = newUser.FindByID(newID)

	companion := new(Companion)
	if err := db.First(&newUser, newID).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}

	companion.UserID = newID
	newUser.Companion = *companion
	db.Save(&companion)
	db.Save(&newUser)
	return companion, nil
}

// ChangeUser changes the user associated with the companion.
func (c *Companion) ChangeUser(newID uint) *Companion {
	db := db.GetDb()
	c.UserID = newID
	db.Save(c)
	return c
}
