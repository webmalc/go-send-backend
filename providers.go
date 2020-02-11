package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/webmalc/go-send-backend/config"
	"github.com/webmalc/go-send-backend/utils"
)

// Returns the controller object
func NewController(manager *DirManager, conf *config.Config) *Controller {
	return &Controller{
		Manager: manager,
		Config:  conf,
	}
}

// Returns the controller object
func NewManager(conf *config.Config) *DirManager {
	return &DirManager{
		Db:     NewRedis(conf),
		Logger: NewLogger(),
		Config: conf,
	}
}

// Setup and get the logger
func NewLogger() *log.Logger {
	logPath := "logs/server.log"
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	utils.ProcessFatalError(err)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	logger := log.New(gin.DefaultWriter, "[APP] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	return logger
}

// Setup and get Redis
func NewRedis(configuration *config.Config) *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     configuration.Database.Host,
		Password: configuration.Database.Password,
		DB:       configuration.Database.Db,
	})

	_, err := db.Ping().Result()
	utils.ProcessFatalError(err)

	return db
}
