package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler serves the index route
func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello!  There's not much to see here :) Please grab one of our front-end samples for use with this sample resource server"})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}

func WhoAmIHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Super secret message!"})
}
