package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return a new Dir structure
func TestGetDir(t *testing.T) {
	path := "/test/path"
	dir := getDir(path)
	assert.Equal(t, dir.Path, path)
}

// Should return a slice with Dir structures
func TestConstructDirsSlice(t *testing.T) {
	dirs := constructDirsSlice([]string{"one", "two"})
	assert.Len(t, dirs, 2)
}

// Should toggle Dir hash
func TestToggleDirHash(t *testing.T) {
	testDelEntry()
	defer testDelEntry()
	dir, err := toggleDirHash(testDirToZip)
	assert := assert.New(t)

	assert.Nil(err)
	assert.NotEmpty(dir.Hash)

	dir, err = toggleDirHash(testDirToZip)

	assert.Nil(err)
	assert.Empty(dir.Hash)

	old := configuration.ZipPath
	configuration.ZipPath = "/invalid/path/"
	defer func() {
		configuration.ZipPath = old
	}()
	_, err = toggleDirHash(testDirToZip)
	assert.NotNil(err)
}

// Should get a Dir structure by hash
func TestGetDirByHash(t *testing.T) {
	dir, err := GetDirByHash(testHash, testPathEncoded)
	assert.Nil(t, err)
	assert.Equal(t, dir.Hash, testHash)
	assert.Equal(t, dir.URL, testExpectedURL)
}

// Should raise errors when the invalid arguments provided
func TestGetDirByHashErrors(t *testing.T) {
	assert := assert.New(t)
	_, err := GetDirByHash(testHash, "invalid_base")
	assert.NotNil(err)

	_, err = GetDirByHash("invalid_hash", testPathEncoded)

	assert.NotNil(err)
	_, _ = db.Del(testPath).Result()
	defer func() {
		_, _ = db.Set(testPath, testHash, 0).Result()
	}()
	_, err = GetDirByHash(testHash, testPathEncoded)

	assert.NotNil(err)
}

func TestGenerateZipErrors(t *testing.T) {
	dir := Dir{Hash: "", Path: ""}
	_, err := generateZip(&dir)
	assert.NotNil(t, err)
}
