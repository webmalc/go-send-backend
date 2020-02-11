package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/config"
	"github.com/webmalc/go-send-backend/files"
)

// Main controller
type Controller struct {
	Manager *DirManager
	Config  *config.Config
}

// Checks the provided error and aborts a controller execution
func checkErrorAndAbort(err error, context *gin.Context) *gin.Error {
	if err != nil {
		return context.AbortWithError(http.StatusBadRequest, err)
	}
	return nil
}

// browseHandler is a handler for listing directories
func (contr *Controller) browseHandler() gin.HandlerFunc {
	var browseHandler gin.HandlerFunc = func(context *gin.Context) {
		path := context.Query("path")
		dirs, err := files.GetDirectories(contr.Config.BasePath, path)
		if checkErrorAndAbort(err, context) != nil {
			return
		}
		context.JSON(http.StatusOK, contr.Manager.constructDirsSlice(dirs))
	}
	return browseHandler
}

// shareHandler is a handler for generating hash for the directory
func (contr *Controller) shareHandler() gin.HandlerFunc {
	var shareHandler gin.HandlerFunc = func(context *gin.Context) {
		path := context.Query("path")
		dir, err := files.ConstructPath(contr.Config.BasePath, path)
		if checkErrorAndAbort(err, context) != nil {
			return
		}
		dirStruct, err := contr.Manager.toggleDirHash(dir)
		if checkErrorAndAbort(err, context) != nil {
			return
		}
		context.JSON(http.StatusOK, dirStruct)
	}
	return shareHandler
}

// getDirectoryHandler is a handler for getting directories
func (contr *Controller) getDirectoryHandler() gin.HandlerFunc {
	var getDirectoryHandler gin.HandlerFunc = func(context *gin.Context) {
		hash := context.Param("hash")
		base := context.Param("base")
		dir, err := contr.Manager.GetDirByHash(hash, base)
		if checkErrorAndAbort(err, context) != nil {
			return
		}
		zip, err := dir.generateZip()
		if checkErrorAndAbort(err, context) != nil {
			return
		}
		context.Header("Content-Description", "File Transfer")
		context.Header("Content-Transfer-Encoding", "binary")
		context.Header("Content-Disposition", "attachment; filename=photos.zip")
		context.Header("Content-Type", "application/octet-stream")
		context.File(zip)
	}
	return getDirectoryHandler
}
