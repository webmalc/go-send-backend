package main

import (
	"github.com/gin-gonic/gin"
)

// Defines the protected routes
func setProtectedRoutes(
	router *gin.RouterGroup,
	controller *Controller,
) *gin.RouterGroup {
	router.GET("/", controller.browseHandler())
	router.GET("/share", controller.shareHandler())
	return router
}

// Defines the public routes
func setPublicRoutes(
	router *gin.RouterGroup,
	controller *Controller,
) *gin.RouterGroup {
	router.GET("/get/:hash/:base", controller.getDirectoryHandler())
	return router
}
