package main

import (
	"github.com/go-redis/redis/v7"
	"github.com/webmalc/go-send-backend/config"
	"github.com/webmalc/go-send-backend/utils"
)

var configuration config.Config
var db *redis.Client

func init() {
	configuration = config.GetConfig()
	db = redis.NewClient(&redis.Options{
		Addr:     configuration.Database.Host,
		Password: configuration.Database.Password,
		DB:       configuration.Database.Db,
	})

	_, err := db.Ping().Result()
	utils.ProcessFatalError(err)

}

func main() {
	runServer(&configuration)
}
