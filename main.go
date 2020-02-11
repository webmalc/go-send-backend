package main

import (
	"github.com/webmalc/go-send-backend/config"
)

// Run the script
func main() {
	configuration := config.NewConfig()
	manager := NewManager(&configuration)
	runServer(manager, &configuration)
}
