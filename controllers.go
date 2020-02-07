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
		context.JSON(http.StatusOK, constructDirsSlice(dirs))
	}
}

// shareHandler is a handler for generating hash for the directory
var shareHandler gin.HandlerFunc = func(context *gin.Context) {
	path := context.Query("path")
	dir, err := files.ConstructPath(configuration.BasePath, path)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	} else {
		dirStruct, err := toggleDirHash(dir)
		if err != nil {
			context.JSON(http.StatusBadRequest, nil)
		} else {
			context.JSON(http.StatusOK, dirStruct)
		}
	}
}

// getDirectoryHandler is a handler for getting directories
var getDirectoryHandler gin.HandlerFunc = func(context *gin.Context) {
	hash := context.Param("hash")
	base := context.Param("base")
	dir, err := GetDirByHash(hash, base)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		zip, err := generateZip(&dir)
		if err != nil {
			context.AbortWithStatus(http.StatusNotFound)
		}
		context.Header("Content-Description", "File Transfer")
		context.Header("Content-Transfer-Encoding", "binary")
		context.Header("Content-Disposition", "attachment; filename=photos.zip")
		context.Header("Content-Type", "application/octet-stream")
		context.File(zip)
	}
}
