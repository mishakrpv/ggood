package tcp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "active"})
}

func signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You have been successfully signed up"})
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You have been successfully logged in"})
}

func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You have been successfully logged out"})
}

func me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Yes it's me"})
}
