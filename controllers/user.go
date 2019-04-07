package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kkjasoncheung/better-buddy-api/models"
)

// UserController is a controller struct for user resource. Contains route handlers.
type UserController struct{}

var userModel = new(models.User)

// TODO: Implement Index, Show, Edit, Destroy actions.

// RetrieveByID handles GET /user/:id. Finds a user by ID.
func (u UserController) RetrieveByID(c *gin.Context) {
	if c.Param("id") != "" {
		// Find the user and return
	}
}

// Retrieve handles GET /user. Returns all users.
func (u UserController) Retrieve(c *gin.Context) {
	// Retrieve all records for user.
}
