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
			user, err := userModel.FindByID(uint(newVal))
			if err != nil {
				if err.Error() == errors.UserNotFoundErrMsg {
					c.JSON(http.StatusNotFound, gin.H{"code": errors.UserNotFoundErrCode, "message": err.Error(), "error": err})
				}
			}
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
	fields := defineFieldsMap(c)
	fmt.Println(fields)
	if user, err := userModel.CreateUser(fields); err == nil {
		c.JSON(http.StatusCreated, gin.H{"user": user})
	} else {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": nil, "message": "Error creating user", "error": err})
		c.Abort()
	}
}

// Update handles PATCH /user/:id. Updates user by ID.
func (u UserController) Update(c *gin.Context) {
	if c.Param("id") == "" {
		err := errors.NewBadRequestError()
		c.JSON(http.StatusBadRequest, gin.H{"code": err.Code, "message": err.Message})
		c.Abort()
		return
	}
	newVal, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err == nil {
		fields := defineFieldsMap(c)
		user, err := userModel.UpdateByID(uint(newVal), fields)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"user": user})
			return
		}
	}
	log.Println(err)
	if err.Error() == errors.UserNotFoundErrMsg {
		c.JSON(http.StatusNotFound, gin.H{"code": errors.UserNotFoundErrCode, "message": err.Error(), "error": err})
	}
	c.JSON(http.StatusInternalServerError, gin.H{"code": nil, "message": "Error to retrieve user", "error": err})
	c.Abort()
	return
}

// Delete handles DELETE /user/:id. Deletes user by ID.
func (u UserController) Delete(c *gin.Context) {
	user := new(models.User)
	if newVal, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		user, err := userModel.DeleteByID(uint(newVal))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": nil, "message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
	} else {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": nil, "message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	return

}

// defineFieldsMap returns a map[string]string with user fields given a context.
func defineFieldsMap(c *gin.Context) map[string]string {
	fields := make(map[string]string)
	if c.PostForm("first_name") != "" {
		fields["first_name"] = c.PostForm("first_name")
	}
	if c.PostForm("last_name") != "" {
		fields["last_name"] = c.PostForm("last_name")
	}
	if c.PostForm("username") != "" {
		fields["username"] = c.PostForm("username")
	}
	if c.PostForm("password") != "" {
		fields["password"] = c.PostForm("password")
	}
	if c.PostForm("email") != "" {
		fields["email"] = c.PostForm("email")
	}
	if c.PostForm("birthday") != "" {
		fields["birthday"] = c.PostForm("birthday")
	}
	if c.PostForm("gender") != "" {
		fields["gender"] = c.PostForm("gender")
	}
	if c.PostForm("display_photo_url") != "" {
		fields["display_photo_url"] = c.PostForm("display_photo_url")
	}

	return fields
}
