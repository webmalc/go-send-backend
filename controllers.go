package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/files"
)

// browseHandler is a handler for listing directories
var browseHandler gin.HandlerFunc = func(context *gin.Context) {
	path := context.Query("path")
	dirs, err := files.GetDirectories(configuration.BasePath, path)
	if err != nil {
		context.JSON(http.StatusBadRequest, dirs)
	} else {
		context.JSON(http.StatusOK, dirs)
	}
}

// getDirectoryHandler is a handler for getting directories
var getDirectoryHandler gin.HandlerFunc = func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"link": "test"})
}
