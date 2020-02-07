package main

import (
	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/config"
	"github.com/webmalc/go-send-backend/utils"
)

// runServer runs server
func runServer(config *config.Config) {
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

	err := router.Run(config.Server)
	utils.ProcessFatalError(err)
}
