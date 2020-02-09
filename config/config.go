package config

import (
	"os"

	"github.com/spf13/viper"
	"github.com/webmalc/go-send-backend/utils"
)

// Database is the database configuration struct
type Database struct {
	Host     string
	Password string
	Db       int
}

// User is the user configuration struct
type User struct {
	Username string
	Password string
}

// Config is the main configuration struct
type Config struct {
	BasePath string
	ZipPath  string
	Host     string
	Prod     bool
	Server   string
	Database Database
	User     User
}

// Return the configuration name
func getConfigName() string {
	env := os.Getenv("GOENV")
	if env != "" {
		return env
	}
	return "main"
}

// GetConfig return the main configuration structure
func GetConfig() Config {
	var config Config

	viper.SetConfigName(getConfigName())
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	utils.ProcessFatalError(err)

	err = viper.Unmarshal(&config)
	utils.ProcessFatalError(err)

	return config

}
