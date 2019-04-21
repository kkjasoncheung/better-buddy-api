package helpers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkjasoncheung/better-buddy-api/errors"
)

// HandleRetrievingUserError Returns a Internal Service Error status code with message.
func HandleRetrievingUserError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(http.StatusInternalServerError, gin.H{"code": nil, "message": "Error to retrieve user", "error": err})
	c.Abort()
}

// HandleBadRequestError Returns a Bad Request Error status code with message.
func HandleBadRequestError(c *gin.Context) {
	err := errors.NewBadRequestError()
	log.Println(err)
	c.JSON(http.StatusBadRequest, gin.H{"code": err.Code, "message": err.Message, "error": err})
	c.Abort()
}
