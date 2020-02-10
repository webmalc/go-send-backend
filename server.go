package main

import (
	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/utils"
)

func setupRouter() *gin.Engine {
	if configuration.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	protectedRouter := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		configuration.User.Username: configuration.User.Password,
	}))
	publicRouter := router.Group("/public")

	setProtectedRoutes(protectedRouter)
	setPublicRoutes(publicRouter)

	return router
}

// Configures and runs the HTTP server
func runServer() {
	router := setupRouter()
	err := router.Run(configuration.Server)
	utils.ProcessFatalError(err)
}
