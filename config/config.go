package config

import (
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
	MaxLevel int
	Prod     bool
	Server   string
	Database Database
	User     User
}

// GetConfig return the main configuration structure
func GetConfig() Config {
	var config Config
	viper.SetConfigName("main")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	utils.ProcessFatalError(err)

	err = viper.Unmarshal(&config)
	utils.ProcessFatalError(err)

	return config

}
