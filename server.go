package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/webmalc/go-send-backend/utils"
)

// Setups the loggers
func setupLoggers() {
	logPath := "logs/server.log"
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	utils.ProcessFatalError(err)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	logger = log.New(gin.DefaultWriter, "[APP] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

// Setups the router
func setupRouter() *gin.Engine {
	if configuration.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	setupLoggers()
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
