package tcp

import (
	"net/http"

	"github.com/ggood/ggood/pkg/logs"
	"github.com/ggood/ggood/pkg/server/di"
	"github.com/ggood/ggood/pkg/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetHTTPHandler(services *di.ServiceContainer) http.Handler {
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
