package utils

import (
	"crypto/rand"
	"fmt"
)

// Checks fatal errors
func ProcessFatalError(err error) {
	if err != nil {
		panic(fmt.Errorf("error: %s", err))
	}
}

// Generates UUID
func GeneateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// Generates a random string
func GenerateRandomString(lenght int) string {
	result := ""
	for i := 0; i < lenght; i++ {
		result += GeneateUUID()
	}
	return result
}
