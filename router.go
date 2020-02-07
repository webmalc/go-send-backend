package main

import (
	"github.com/gin-gonic/gin"
)

// Defines the protected routes
func setProtectedRoutes(router *gin.RouterGroup) *gin.RouterGroup {

	router.GET("/", browseHandler)
	return router
}

// Defines the public routes
func setPublicRoutes(router *gin.RouterGroup) *gin.RouterGroup {

	router.GET("/get", getDirectoryHandler)
	return router
}
