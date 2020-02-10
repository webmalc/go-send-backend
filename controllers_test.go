package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GOENV", "test")
	configuration.BasePath = testWorkingPath
}

// Returns the request with basic auth
func getAdminRequest(url string) *http.Request {
	req := getRequest(url)
	req.SetBasicAuth(configuration.User.Username, configuration.User.Password)
	return req
}

// Returns the request
func getRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req
}

// Should return the 401 code for the unauthorized request
func TestAuth(t *testing.T) {
	router := setupRouter(&configuration)

	request := getRequest("/admin/")
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	assert.Equal(t, writer.Code, 401)

	router = setupRouter(&configuration)
	request.SetBasicAuth("invalid", "invalid")
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	assert.Equal(t, writer.Code, 401)
}

// Should return the JSON with directories
func TestBrowseHandler(t *testing.T) {
	router := setupRouter(&configuration)

	request := getAdminRequest("/admin/")
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	assert.Equal(t, writer.Code, 200)

	var data []Dir
	err := json.Unmarshal(writer.Body.Bytes(), &data)

	assert.Nil(t, err)
	assert.Contains(t, data, Dir{Path: "config/"})
	assert.Contains(t, data, Dir{Path: "utils/"})
}

// Should return an error with the invalid path
func TestBrowseHandlerError(t *testing.T) {
	router := setupRouter(&configuration)

	request := getAdminRequest("/admin/")
	q := request.URL.Query()
	q.Add("path", "/invalid/path")
	request.URL.RawQuery = q.Encode()

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	assert.Equal(t, writer.Code, 400)
}

// Should return the JSON with Dir structure
func TestShareHandler(t *testing.T) {
	router := setupRouter(&configuration)
	configuration.BasePath = testWorkingPath

	request := getAdminRequest("/admin/share")
	q := request.URL.Query()
	q.Add("path", testDirToZip)
	request.URL.RawQuery = q.Encode()

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	assert.Equal(t, writer.Code, 200)

	var dir Dir
	err := json.Unmarshal(writer.Body.Bytes(), &dir)

	assert.Nil(t, err)
	assert.NotEmpty(t, dir.URL)

	_, _ = dir.toggleHash()
}

// Should return an error with the invalid path
func TestShareHandlerErrors(t *testing.T) {
	router := setupRouter(&configuration)

	request := getAdminRequest("/admin/share")
	q := request.URL.Query()
	q.Add("path", "/root/")
	request.URL.RawQuery = q.Encode()

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	assert.Equal(t, writer.Code, 400)
}

// Should return a response with the zip archive
func TestGetDirectoryHandler(t *testing.T) {
	path := testPath + "utils"
	PathEncoded := base64.StdEncoding.EncodeToString([]byte(path))
	configuration.BasePath = testWorkingPath
	err := db.Set(path, testHash, 0).Err()
	if err != nil {
		panic(err)
	}
	router := setupRouter(&configuration)

	url := fmt.Sprintf("public/get/%s/%s", testHash, PathEncoded)
	request := getRequest(url)

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	assert.Equal(t, writer.Code, 200)
	assert.Equal(t, writer.Header().Get("Content-Description"), "File Transfer")

	dir := Dir{Path: path, Hash: testHash}
	_, _ = dir.removeHash()

}

// Should return an error with the invalid path
func TestGetDirectoryHandlerErrors(t *testing.T) {
	router := setupRouter(&configuration)
	request := getRequest("public/get/:hash/:base")

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	assert.Equal(t, writer.Code, 400)
}
