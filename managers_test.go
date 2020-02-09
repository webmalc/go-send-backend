package main

import (
	"testing"
)

// Should return a new Dir structure
func TestGetDir(t *testing.T) {
	path := "/test/path"
	dir := getDir(path)
	if dir.Path != path {
		t.Errorf("Path should been %s, got %s", path, dir.Path)
	}
}

// Should return a slice with Dir structures
func TestConstructDirsSlice(t *testing.T) {
	dirs := constructDirsSlice([]string{"one", "two"})

	if len(dirs) != 2 {
		t.Errorf("two Dir structures should have been returned")
	}
}

// Should toggle Dir hash
func TestToggleDirHash(t *testing.T) {
	testDelEntry()
	defer testDelEntry()
	dir, err := toggleDirHash(testDirToZip)
	if err != nil {
		t.Errorf("error %s", err)
	}
	if dir.Hash == "" {
		t.Errorf("hash has not been set")
	}
	dir, err = toggleDirHash(testDirToZip)
	if err != nil {
		t.Errorf("error %s", err)
	}
	if dir.Hash != "" {
		t.Errorf("hash has not been deleted")
	}
	old := configuration.ZipPath
	configuration.ZipPath = "/invalid/path/"
	defer func() {
		configuration.ZipPath = old
	}()
	_, err = toggleDirHash(testDirToZip)
	if err == nil {
		t.Errorf("the zip file should not have been created")
	}
}

// Should get a Dir structure by hash
func TestGetDirByHash(t *testing.T) {
	dir, err := GetDirByHash(testHash, testPathEncoded)
	if err != nil {
		t.Errorf("error %s", err)
	}
	if dir.Hash != testHash {
		t.Errorf("Got the invalid hash %s, expected %s", dir.Hash, testHash)
	}
	if dir.URL != testExpectedURL {
		t.Errorf("Got the invalid URL %s, expected %s", dir.URL, testExpectedURL)
	}
}

// Should raise errors when the invalid arguments provided
func TestGetDirByHashErrors(t *testing.T) {
	_, err := GetDirByHash(testHash, "invalid_base")
	if err == nil {
		t.Errorf("Base should not have been decoded")
	}
	_, err = GetDirByHash("invalid_hash", testPathEncoded)

	if err == nil {
		t.Errorf("Hashes must not been equal")
	}
	_, _ = db.Del(testPath).Result()
	defer func() {
		_, _ = db.Set(testPath, testHash, 0).Result()
	}()
	_, err = GetDirByHash(testHash, testPathEncoded)

	if err == nil {
		t.Errorf("Base should not have been found in the DB")
	}
}

func TestGenerateZipErrors(t *testing.T) {
	dir := Dir{Hash: "", Path: ""}
	_, err := generateZip(&dir)
	if err == nil {
		t.Errorf("The zip archive should not have been generated")
	}
}
