package server

import (
	"net/http"

	"github.com/ggood/ggood/pkg/logs"
	"github.com/ggood/ggood/pkg/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func addRoutes() http.Handler {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(logs.NoLevel(log.Logger, zerolog.InfoLevel)))

	authGroup := r.Group("/auth")

	authGroup.POST("/signup", signup)
	authGroup.POST("/login", login)
	authGroup.POST("/logout", logout)

	r.GET("/me", middleware.Authorize(me))

	r.GET("/health", health)

	return r
}

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
