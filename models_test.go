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
	assertT := assert.New(t)

	zip := configuration.ZipPath

	dir := Dir{Path: testDirToZip}
	err := dir.toggleHash()

	assertT.Nil(err)
	assertT.NotEmpty(dir.Hash)
	assertT.NotEmpty(dir.URL)

	zip += dir.Hash + ".zip"
	_, err = os.Stat(zip)
	assertT.Nil(err)

	err = dir.toggleHash()

	assertT.Nil(err)
	assertT.Empty(dir.Hash)
	assertT.Empty(dir.URL)

	_, err = os.Stat(zip)
	assertT.NotNil(err)

	old := configuration.ZipPath
	configuration.ZipPath = "/invalid/path/"
	defer func() {
		configuration.ZipPath = old
	}()
	err = dir.toggleHash()
	assertT.NotNil(err)
}
