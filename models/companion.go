package models

import (
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

// GetUser returns the user associated with this companion.
func (c Companion) GetUser() *User {
	db := db.GetDb()
	user := new(User)
	db.First(&user, c.UserID)
	return user
}

// ChangeUser changes the user associated with the companion.
func (c *Companion) ChangeUser(newID uint) *Companion {
	db := db.GetDb()
	c.UserID = newID
	db.Save(c)
	return c
}

// GetAllCompanions gets all companions in database.
func (c Companion) GetAllCompanions() []Companion {
	db := db.GetDb()
	companions := []Companion{}
	db.Find(&companions)
	return companions
}
