package main

import (
	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/config"
	"github.com/webmalc/go-send-backend/utils"
)

func setupRouter(config *config.Config) *gin.Engine {
	if config.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	protected_router := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		configuration.User.Username: configuration.User.Password,
	}))
	public_router := router.Group("/public")

	setProtectedRoutes(protected_router)
	setPublicRoutes(public_router)

	return router
}

// Configures and runs the HTTP server
func runServer(config *config.Config) {
	router := setupRouter(config)
	err := router.Run(config.Server)
	utils.ProcessFatalError(err)
}
