package main

import (
	"encoding/base64"
	"errors"

	"github.com/webmalc/go-send-backend/files"
)

// Returns a new constructed Dir object
func getDir(dir string) Dir {
	dirStruct := Dir{}
	dirStruct.constructor(dir)
	return dirStruct
}

// Constructs the Dir slice from paths
func constructDirsSlice(dirs []string) []Dir {
	var results []Dir
	for _, dir := range dirs {
		results = append(results, getDir(dir))
	}
	return results
}

// Toggles hash for the directory
func toggleDirHash(dir string) (Dir, error) {
	dirStruct := getDir(dir)
	err := dirStruct.toggleHash()
	if err != nil {
		return dirStruct, err
	}
	return dirStruct, nil
}

// Gets a Dir structure by the hash
func GetDirByHash(hash, base string) (Dir, error) {
	dir := Dir{}
	decoded, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return dir, err
	}
	key := string(decoded)
	dbHash, err := db.Get(key).Result()
	if err != nil {
		return dir, err
	}
	if dbHash != hash {
		return dir, errors.New("unable to find the directory")
	}
	dir.Path = key
	dir.Hash = hash
	dir.setURL()

	return dir, nil
}

// Generates a zip archive for the Dir object
func generateZip(dir *Dir) (string, error) {
	if dir.Path != "" && dir.Hash != "" {
		return files.ZipDir(dir.Path, configuration.ZipPath, dir.Hash)
	}
	return "", errors.New("unable to generate a zip archive")
}
