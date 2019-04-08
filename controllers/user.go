package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kkjasoncheung/better-buddy-api/models"
)

// UserController is a controller struct for user resource. Contains route handlers.
type UserController struct{}

var userModel = new(models.User)

// Retrieve handles GET /user. Returns all users.
func (u UserController) Retrieve(c *gin.Context) {
	// Retrieve all records for user.
	users := userModel.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// RetrieveByID handles GET /user/:id. Finds a user by ID.
func (u UserController) RetrieveByID(c *gin.Context) {
	if c.Param("id") != "" {
		// Find the user and return
		if newVal, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
			user := userModel.FindByID(uint(newVal))
			c.JSON(http.StatusOK, gin.H{"user": user})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

// TODO: Implement PATCH user/:id, POST user, DELETE user/:id
