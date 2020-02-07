package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/webmalc/go-send-backend/utils"
)

// Dir is the directory struct
type Dir struct {
	Path string
	Hash string
	URL  string
}

// constructor gets the Dir
func (dirStruct *Dir) constructor(dir string) *Dir {
	dirStruct.Path = dir
	dirStruct.setHashFromDB()
	dirStruct.setURL()

	return dirStruct
}

// getHash gets the Dir hash from the database
func (dir *Dir) getHash() string {
	hash, _ := db.Get(dir.Path).Result()
	return hash
}

// setHashFromDB sets the Dir hash from the database
func (dir *Dir) setHashFromDB() *Dir {
	hash, _ := db.Get(dir.Path).Result()
	dir.Hash = hash
	return dir
}

// setURL sets URL for the Dir
func (dir *Dir) setURL() *Dir {
	dir.URL = ""
	if dir.Hash != "" && dir.Path != "" {
		pattern := "%s/public/get/%s/%s"
		host := configuration.Host
		base := base64.StdEncoding.EncodeToString([]byte(dir.Path))
		dir.URL = fmt.Sprintf(pattern, host, dir.Hash, base)
	}
	return dir
}

// generateZip generates zip for the Dir object
func generateZip(dir *Dir) (string, error) {
	if dir.Path != "" && dir.Hash != "" {
		return utils.ZipDir(dir.Path, configuration.ZipPath, dir.Hash)
	}
	return "", errors.New("unable to generate a zip archive")
}

// toggleHash toggles the Dir hash
func (dirStruct *Dir) toggleHash() (*Dir, error) {
	if dirStruct.Hash == "" {
		hash := utils.GenerateRandomString(5)
		err := db.Set(dirStruct.Path, hash, 0).Err()
		if err != nil {
			return dirStruct, err
		}
		dirStruct.Hash = hash
		_, err = generateZip(dirStruct)
		if err != nil {
			return dirStruct, err
		}
	} else {
		_, err := db.Del(dirStruct.Path).Result()
		if err != nil {
			return dirStruct, err
		}
		utils.DeleteFile(configuration.ZipPath, dirStruct.Hash)
		dirStruct.Hash = ""
	}
	dirStruct.setURL()
	return dirStruct, nil
}

// GetDirByHash gets a Dir structure by the hash
func GetDirByHash(hash string, base string) (Dir, error) {
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

// constructDirsSlice constructs the Dir slice from paths
func constructDirsSlice(dirs []string) []Dir {
	var results []Dir
	for _, dir := range dirs {
		dirStruct := Dir{}
		dirStruct.constructor(dir)
		results = append(results, dirStruct)
	}
	return results
}

// toggleDirHash toggles hash for the directory
func toggleDirHash(dir string) (Dir, error) {
	dirStruct := Dir{}
	dirStruct.constructor(dir)
	_, err := dirStruct.toggleHash()
	if err != nil {
		return dirStruct, err
	}
	return dirStruct, nil
}
