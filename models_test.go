package main

import (
	"os"
	"testing"
)

// Should construct a Dir structure
func TestDir_constructor(t *testing.T) {
	dir := Dir{}
	dir.constructor(testPath)
	if dir.Hash != testHash {
		t.Errorf("Expected hash %s, got %s", testHash, dir.Hash)
	}
	if dir.URL != testExpectedURL {
		t.Errorf("Expected URL %s, got %s", testExpectedURL, dir.URL)
	}
}

// Should get a hash from the database
func TestDir_setHashFromDB(t *testing.T) {
	dir := Dir{Path: testPath}
	dir.setHashFromDB()
	if dir.Hash != testHash {
		t.Errorf("Expected hash %s, got %s", testHash, dir.Hash)
	}
}

// Should compose the URL to get the zip archive
func TestDir_setURL(t *testing.T) {
	dir := Dir{Path: testPath, Hash: testHash}
	dir.setURL()
	if dir.URL != testExpectedURL {
		t.Errorf("Expected URL %s, got %s", testExpectedURL, dir.URL)
	}
}

// Should toggle a Dir structure hash
func TestDir_toggleHash(t *testing.T) {
	testDelEntry()
	defer testDelEntry()

	zip := configuration.ZipPath

	dir := Dir{Path: testDirToZip}
	_, err := dir.toggleHash()
	if err != nil {
		t.Errorf("error %s", err)
	}
	if dir.Hash == "" {
		t.Errorf("hash has not been set")
	}
	if dir.URL == "" {
		t.Errorf("URL has not been set")
	}
	zip += dir.Hash + ".zip"
	if _, err := os.Stat(zip); err != nil {
		t.Errorf("the zip file has not been created. file: %s", zip)
	}
	_, err = dir.toggleHash()
	if err != nil {
		t.Errorf("error %s", err)
	}
	if dir.Hash != "" {
		t.Errorf("hash has not been deleted")
	}
	if dir.URL != "" {
		t.Errorf("URL has not been deleted")
	}
	if _, err := os.Stat(zip); err == nil {
		t.Errorf("the zip file has not been deleted. file: %s", zip)
	}
	old := configuration.ZipPath
	configuration.ZipPath = "/invalid/path/"
	defer func() {
		configuration.ZipPath = old
	}()
	_, err = dir.toggleHash()
	if err == nil {
		t.Errorf("the zip file should not have been created")
	}
}
