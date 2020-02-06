package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Database is the database configuration struct
type Database struct {
	Username string
	Password string
	Host     string
	Port     int
}

// Config is the main configuration struct
type Config struct {
	BasePath string
	MaxLevel int
	Database Database
}

// GetConfig return the main configuration structure
func GetConfig() Config {
	var config Config
	viper.SetConfigName("main")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error configuration file: %s", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %s", err))
	}
	return config

}
