package main

import (
	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/config"
)

// Defines the protected routes
func setProtectedRoutes(
	router *gin.RouterGroup,
	manager *DirManager,
	conf *config.Config,
) *gin.RouterGroup {
	router.GET("/", browseHandler(manager, conf))
	router.GET("/share", shareHandler(manager, conf))
	return router
}

// Defines the public routes
func setPublicRoutes(
	router *gin.RouterGroup,
	manager *DirManager,
) *gin.RouterGroup {
	router.GET("/get/:hash/:base", getDirectoryHandler(manager))
	return router
}
