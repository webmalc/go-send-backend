package main

import (
	"encoding/base64"
	"os"
	"testing"
)

var testWorkingPath string
var testPath string
var testPathEncoded string
var testHash string
var testExpectedURL string
var testDirToZip string

// Remove the Dir entry from the DB
func testDelEntry() {
	_, err := db.Del(testDirToZip).Result()
	if err != nil {
		panic(err)
	}
}

func testSetUp() {
	os.Setenv("GOENV", "test")
	testWorkingPath, _ = os.Getwd()
	testPath = testWorkingPath + "/"
	testPathEncoded = base64.StdEncoding.EncodeToString([]byte(testPath))
	testHash = "testhash"
	testExpectedURL = configuration.Host + "/public/get/"
	testExpectedURL += testHash + "/" + testPathEncoded
	testDirToZip = testWorkingPath + "/utils/"
	configuration.ZipPath = testWorkingPath + "/"
	err := db.Set(testPath, testHash, 0).Err()
	if err != nil {
		panic(err)
	}
}

// Setup the package
func TestMain(m *testing.M) {
	testSetUp()
	os.Exit(m.Run())
	_, err := db.Del(testPath).Result()
	if err != nil {
		panic(err)
	}
}
