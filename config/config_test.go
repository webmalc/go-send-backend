package config

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("GOENV", "test")
}

// Should return the config name based on the environment variable GOENV
func TestGetConfigName(t *testing.T) {
	if name := getConfigName(); name != "test" {
		t.Errorf("name should have been 'test'. Got: %s", name)
	}
	os.Setenv("GOENV", "abc")
	defer func() {
		os.Setenv("GOENV", "test")
	}()

	if name := getConfigName(); name != "abc" {
		t.Errorf("name should have been 'abc'. Got: %s", name)
	}

	os.Setenv("GOENV", "")
	if name := getConfigName(); name != "main" {
		t.Errorf("name should have been 'main'. Got: %s", name)
	}

}

// Should return a configuration object
func TestGetConfig(t *testing.T) {
	config := GetConfig()
	basePath := "/path/to/directories"
	db := 9
	user := "user"
	password := "password"

	if config.BasePath != basePath {
		t.Errorf("Value should be %s", basePath)
	}
	if config.Database.Db != db {
		t.Errorf("Value should be %d", db)
	}
	if config.User.Username != user {
		t.Errorf("Value should be %s", user)
	}
	if config.User.Password != password {
		t.Errorf("Value should be %s", password)
	}
}
