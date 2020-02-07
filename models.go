package main

import (
	"github.com/webmalc/go-send-backend/utils"
)

// Dir is the directory struct
type Dir struct {
	Path string
	Hash string
}

// constructor gets the Dir
func (dirStruct *Dir) constructor(dir string) *Dir {
	dirStruct.Path = dir
	dirStruct.setHashFromDB()

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

// toggleHash toggles the Dir hash
func (dirStruct *Dir) toggleHash() (*Dir, error) {
	if dirStruct.Hash == "" {
		hash := utils.GenerateRandomString(5)
		err := db.Set(dirStruct.Path, hash, 0).Err()
		if err != nil {
			return dirStruct, err
		}
		dirStruct.Hash = hash
	} else {
		_, err := db.Del(dirStruct.Path).Result()
		if err != nil {
			return dirStruct, err
		}
		dirStruct.Hash = ""
	}
	return dirStruct, nil
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
