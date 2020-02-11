package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return a link to the Controller instance
func TestNewController(t *testing.T) {
	manager := NewManager(&configuration)
	controller := NewController(manager, &configuration)
	assert.Equal(t, "user", controller.Config.User.Username)
}

// Should return a link to the Manager instance
func TestNewManager(t *testing.T) {
	manager := NewManager(&configuration)
	assert.Equal(t, "user", manager.Config.User.Username)
}

// Should return a link to the Logger instance
func TestNewLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger()
	logger.SetOutput(&buf)
	message := "test message"
	logger.Print(message)
	assert.Contains(t, buf.String(), message)
}

// Should return a link to the DB instance
func TestNewRedis(t *testing.T) {
	db := NewRedis(&configuration)
	_, err := db.Ping().Result()
	assert.Nil(t, err)
	db.Close()
}
