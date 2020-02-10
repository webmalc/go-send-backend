package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should construct a Dir structure
func TestDir_constructor(t *testing.T) {
	dir := Dir{}
	dir.constructor(testPath)

	assert.Equal(t, dir.Hash, testHash)
	assert.Equal(t, dir.URL, testExpectedURL)
}

// Should get a hash from the database
func TestDir_setHashFromDB(t *testing.T) {
	dir := Dir{Path: testPath}
	dir.setHashFromDB()

	assert.Equal(t, dir.Hash, testHash)
}

// Should compose the URL to get the zip archive
func TestDir_setURL(t *testing.T) {
	dir := Dir{Path: testPath, Hash: testHash}
	dir.setURL()

	assert.Equal(t, dir.URL, testExpectedURL)
}

// Should toggle a Dir structure hash
func TestDir_toggleHash(t *testing.T) {
	testDelEntry()
	defer testDelEntry()
	assert := assert.New(t)

	zip := configuration.ZipPath

	dir := Dir{Path: testDirToZip}
	_, err := dir.toggleHash()

	assert.Nil(err)
	assert.NotEmpty(dir.Hash)
	assert.NotEmpty(dir.URL)

	zip += dir.Hash + ".zip"
	_, err = os.Stat(zip)
	assert.Nil(err)

	_, err = dir.toggleHash()

	assert.Nil(err)
	assert.Empty(dir.Hash)
	assert.Empty(dir.URL)

	_, err = os.Stat(zip)
	assert.NotNil(err)

	old := configuration.ZipPath
	configuration.ZipPath = "/invalid/path/"
	defer func() {
		configuration.ZipPath = old
	}()
	_, err = dir.toggleHash()
	assert.NotNil(err)
}
