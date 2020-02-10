package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return a new Dir structure
func TestGetDir(t *testing.T) {
	path := "/test/path"
	dir := manager.getDir(path)
	assert.Equal(t, dir.Path, path)
}

// Should return a slice with Dir structures
func TestConstructDirsSlice(t *testing.T) {
	dirs := manager.constructDirsSlice([]string{"one", "two"})
	assert.Len(t, dirs, 2)
}

// Should toggle Dir hash
func TestToggleDirHash(t *testing.T) {
	testDelEntry()
	defer testDelEntry()
	dir, err := manager.toggleDirHash(testDirToZip)
	assert.Nil(t, err)
	assert.NotEmpty(t, dir.Hash)

	dir, err = manager.toggleDirHash(testDirToZip)

	assert.Nil(t, err)
	assert.Empty(t, dir.Hash)

	old := configuration.ZipPath
	configuration.ZipPath = "/some/invalid/path/"
	defer func() {
		configuration.ZipPath = old
	}()
	_, err = manager.toggleDirHash(testDirToZip)
	assert.NotNil(t, err)
}

// Should get a Dir structure by hash
func TestGetDirByHash(t *testing.T) {
	dir, err := manager.GetDirByHash(testHash, testPathEncoded)
	assert.Nil(t, err)
	assert.Equal(t, dir.Hash, testHash)
	assert.Equal(t, dir.URL, testExpectedURL)
}

// Should raise errors when the invalid arguments provided
func TestGetDirByHashErrors(t *testing.T) {
	_, err := manager.GetDirByHash(testHash, "invalid_base")
	assert.NotNil(t, err)

	_, err = manager.GetDirByHash("invalid_hash", testPathEncoded)

	assert.NotNil(t, err)
	_, _ = db.Del(testPath).Result()
	defer func() {
		_, _ = db.Set(testPath, testHash, 0).Result()
	}()
	_, err = manager.GetDirByHash(testHash, testPathEncoded)

	assert.NotNil(t, err)
}

func TestGenerateZipErrors(t *testing.T) {
	dir := Dir{Hash: "", Path: ""}
	_, err := dir.generateZip()
	assert.NotNil(t, err)
}
