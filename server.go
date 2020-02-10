package main

import (
	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/config"
	"github.com/webmalc/go-send-backend/utils"
)

// Setups the router
func setupRouter(manager *DirManager, conf *config.Config) *gin.Engine {
	if conf.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	protectedRouter := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		conf.User.Username: conf.User.Password,
	}))
	publicRouter := router.Group("/public")

	setProtectedRoutes(protectedRouter, manager, conf)
	setPublicRoutes(publicRouter, manager)

	return router
}

// Configures and runs the HTTP server
func runServer(manager *DirManager, conf *config.Config) {
	router := setupRouter(manager, conf)
	err := router.Run(conf.Server)
	utils.ProcessFatalError(err)
}
