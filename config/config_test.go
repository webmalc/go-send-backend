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
	assert := assert.New(t)
	assert.Equal(getConfigName(), "test")

	os.Setenv("GOENV", "abc")
	defer func() {
		os.Setenv("GOENV", "test")
	}()
	assert.Equal(getConfigName(), "abc")

	os.Setenv("GOENV", "")
	assert.Equal(getConfigName(), "main")

}

// Should return a configuration object
func TestGetConfig(t *testing.T) {
	assert := assert.New(t)
	config := GetConfig()

	assert.Equal(config.BasePath, "/path/to/directories")
	assert.Equal(config.Database.Db, 9)
	assert.Equal(config.User.Username, "user")
	assert.Equal(config.User.Password, "password")
}
