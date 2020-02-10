package main

import (
	"encoding/base64"
	"os"
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
	configuration = config.GetConfig()
	db = getRedis(&configuration)
	manager = &DirManager{
		Db:     db,
		Logger: getLogger(),
		Config: &configuration,
	}
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
