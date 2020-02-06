package main

import (
	"fmt"
	"github.com/webmalc/go-send-backend/config"
)

func main() {
	config := config.GetConfig()
	fmt.Println(config)
}
