package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kkjasoncheung/better-buddy-api/errors"
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
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": nil, "message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
	}
	err := errors.NewBadRequestError()
	log.Println(err)
	c.JSON(http.StatusBadRequest, gin.H{"code": err.Code, "message": err.Message, "error": err})
	c.Abort()
	return
}

// Create handles POST /user. Creates a new user.
func (u UserController) Create(c *gin.Context) {
	fields := make(map[string]string)
	fields["first_name"] = c.PostForm("first_name")
	fields["last_name"] = c.PostForm("last_name")
	fields["username"] = c.PostForm("username")
	fields["password"] = c.PostForm("password")
	fields["email"] = c.PostForm("email")
	fields["birthday"] = c.PostForm("birthday")
	fields["gender"] = c.PostForm("gender")
	fields["display_photo_url"] = c.PostForm("display_photo_url")
	fmt.Println(fields)
	if user, err := userModel.CreateUser(fields); err == nil {
		c.JSON(http.StatusCreated, gin.H{"user": user})
	} else {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": nil, "message": "Error creating user", "error": err})
		c.Abort()
	}
}

// TODO: Implement PATCH user/:id, DELETE user/:id
