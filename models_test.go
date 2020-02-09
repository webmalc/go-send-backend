package main

import (
	"encoding/base64"
	"os"
	"testing"
)

var path string
var pathEncoded string
var hash string
var expectedURL string

// Initializes the main variables
func init() {
	os.Setenv("GOENV", "test")
	path = "/test/path/"
	pathEncoded = base64.StdEncoding.EncodeToString([]byte(path))
	hash = "testhash"
	expectedURL = configuration.Host + "/public/get/" + hash + "/" + pathEncoded
	err := db.Set(path, hash, 0).Err()
	if err != nil {
		panic(err)
	}
}

// Should construct a Dir structure
func TestDir_constructor(t *testing.T) {
	dir := Dir{}
	dir.constructor(path)
	if dir.Hash != hash {
		t.Errorf("Expected hash %s, got %s", hash, dir.Hash)
	}
	if dir.URL != expectedURL {
		t.Errorf("Expected URL %s, got %s", expectedURL, dir.URL)
	}
}

// Should get a hash from the database
func TestDir_setHashFromDB(t *testing.T) {
	dir := Dir{Path: path}
	dir.setHashFromDB()
	if dir.Hash != hash {
		t.Errorf("Expected hash %s, got %s", hash, dir.Hash)
	}
}

// Should compose the URL to get the zip archive
func TestDir_setURL(t *testing.T) {
	dir := Dir{Path: path, Hash: hash}
	dir.setURL()
	if dir.URL != expectedURL {
		t.Errorf("Expected URL %s, got %s", expectedURL, dir.URL)
	}
}

// Should toggle a Dir structure hash
func TestDir_toggleHash(t *testing.T) {
	workingPath, _ := os.Getwd()
	dirToZip := workingPath + "/utils/"
	configuration.ZipPath = workingPath + "/"
	zip := configuration.ZipPath

	dir := Dir{Path: dirToZip}
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
	configuration.ZipPath = "/invalid/path/"
	_, err = dir.toggleHash()
	if err == nil {
		t.Errorf("the zip file should not have been created")
	}
}
