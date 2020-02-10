package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should generate a random UUID
func TestGeneateUUID(t *testing.T) {
	uuid1 := GeneateUUID()
	uuid2 := GeneateUUID()
	uuidLen1 := len(uuid1)
	uuidLen2 := len(uuid2)

	assert.Equal(t, uuidLen1, 36)
	assert.Equal(t, uuidLen1, uuidLen2)
	assert.NotEqual(t, uuid1, uuid2)
}

// Should panic when randReader failed
func TestGeneateUUIDPanic(t *testing.T) {
	old := randReader
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
		randReader = old
	}()
	randReader = func(b []byte) (n int, err error) {
		return 0, errors.New("test error")
	}
	GeneateUUID()
}

// Should generate a random string
func TestGenerateRandomString(t *testing.T) {
	str1 := GenerateRandomString(3)
	str2 := GenerateRandomString(3)
	strLen1 := len(str1)
	strLen2 := len(str2)

	assert.Equal(t, strLen1, 36*3)
	assert.Equal(t, strLen1, strLen2)
	assert.NotEqual(t, str1, str2)
	assert.Equal(t, len(GenerateRandomString(5)), 36*5)
}

// Should panic if the provided error is not nil
func TestProcessFatalErrorPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ProcessFatalError(errors.New("test error"))
}

// Should do nothing if the provided error is nil
func TestProcessFatalErrorNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic")
		}
	}()
	ProcessFatalError(nil)
}
