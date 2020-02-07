package main

import (
	"github.com/webmalc/go-send-backend/config"
)

var configuration config.Config

func init() {
	configuration = config.GetConfig()
}

func main() {
	runServer(&configuration)
}
