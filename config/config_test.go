package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GOENV", "test")
}

// Should return the config name based on the environment variable GOENV
func TestGetConfigName(t *testing.T) {
	assertT := assert.New(t)
	assertT.Equal(getConfigName(), "test")

	os.Setenv("GOENV", "abc")
	defer func() {
		os.Setenv("GOENV", "test")
	}()
	assertT.Equal(getConfigName(), "abc")

	os.Setenv("GOENV", "")
	assertT.Equal(getConfigName(), "main")
}

// Should return a configuration object
func TestNewConfig(t *testing.T) {
	assertT := assert.New(t)
	config := NewConfig()

	assertT.Equal(config.BasePath, "/path/to/directories")
	assertT.Equal(config.Database.Db, 9)
	assertT.Equal(config.User.Username, "user")
	assertT.Equal(config.User.Password, "password")
}
