package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/config"
	"github.com/webmalc/go-send-backend/utils"
)

// Returns the CORS configuration
func getCors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	return cors.New(corsConfig)
}

// Setups the router
func setupRouter(manager *DirManager, conf *config.Config) *gin.Engine {
	if conf.Prod {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(getCors())
	router.LoadHTMLGlob("templates/*")
	protectedRouter := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		conf.User.Username: conf.User.Password,
	}))
	publicRouter := router.Group("/public")

	controller := NewController(manager, conf)
	setProtectedRoutes(protectedRouter, controller)
	setPublicRoutes(publicRouter, controller)

	return router
}

// Configures and runs the HTTP server
func runServer(manager *DirManager, conf *config.Config) {
	router := setupRouter(manager, conf)
	err := router.Run(conf.Server)
	utils.ProcessFatalError(err)
}
