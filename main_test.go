package main

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-redis/redis/v7"
	"github.com/webmalc/go-send-backend/config"
)

var (
	testWorkingPath string
	testPath        string
	testPathEncoded string
	testHash        string
	testExpectedURL string
	testDirToZip    string
	db              *redis.Client
	configuration   config.Config
	manager         *DirManager
)

// Remove the Dir entry from the DB
func testDelEntry() {
	_, err := db.Del(testDirToZip).Result()
	if err != nil {
		panic(err)
	}
}

func testSetUp() {
	os.Setenv("GOENV", "test")
	configuration = config.NewConfig()
	db = NewRedis(&configuration)
	manager = &DirManager{
		Db:     db,
		Logger: NewLogger(),
		Config: &configuration,
	}
	testWorkingPath, _ = os.Getwd()
	testWorkingPath, _ = filepath.Abs(testWorkingPath)
	testPath = testWorkingPath + "/"
	testPathEncoded = base64.StdEncoding.EncodeToString([]byte(testPath))
	testHash = "testhash"
	testExpectedURL = configuration.Host + "/public/get/"
	testExpectedURL += testHash + "/" + testPathEncoded
	testDirToZip = testWorkingPath + "/utils/"
	configuration.ZipPath = testWorkingPath + "/"
	configuration.BasePath = testWorkingPath + "/"
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
