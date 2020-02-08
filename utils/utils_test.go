package utils

import (
	"errors"
	"testing"
)

// Should generate a random UUID
func TestGeneateUUID(t *testing.T) {
	uuid1 := GeneateUUID()
	uuid2 := GeneateUUID()
	uuidLen1 := len(uuid1)
	uuidLen2 := len(uuid2)
	if uuidLen1 != 36 || uuidLen2 != uuidLen1 {
		t.Errorf("invalid UUIDs length: %d, %d", uuidLen2, uuidLen1)
	}
	if uuid1 == uuid2 {
		t.Errorf("equal UUIDs: %s, %s", uuid1, uuid1)
	}
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

	if strLen1 != 36*3 || strLen1 != strLen2 {
		t.Errorf("invalid strings length: %d, %d", strLen1, strLen2)
	}
	if str1 == str2 {
		t.Errorf("equal strings: %s, %s", str1, str2)
	}
	if strLenLong := len(GenerateRandomString(5)); strLenLong != 36*5 {
		t.Errorf("invalid strings length: %d", strLenLong)
	}
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
