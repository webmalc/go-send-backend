package main

import (
	"encoding/base64"
	"fmt"

	"github.com/webmalc/go-send-backend/files"
	"github.com/webmalc/go-send-backend/utils"
)

// The directory struct
type Dir struct {
	Path string
	Hash string
	URL  string
}

// Constructs a Dir structure
func (dirStruct *Dir) constructor(dir string) *Dir {
	dirStruct.Path = dir
	dirStruct.setHashFromDB()
	dirStruct.setURL()

	return dirStruct
}

// Sets the Dir hash from the database
func (dir *Dir) setHashFromDB() *Dir {
	hash, _ := db.Get(dir.Path).Result()
	dir.Hash = hash
	return dir
}

// Sets URL for the Dir
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

// Removes the Dir hash
func (dirStruct *Dir) removeHash() (*Dir, error) {
	_, err := db.Del(dirStruct.Path).Result()
	if err != nil {
		return dirStruct, err
	}
	_ = files.DeleteZip(configuration.ZipPath, dirStruct.Hash)
	dirStruct.Hash = ""

	return dirStruct, nil
}

// Sets the Dir hash
func (dirStruct *Dir) setHash() (*Dir, error) {

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
	return dirStruct, nil
}

// Toggles the Dir hash
func (dirStruct *Dir) toggleHash() (*Dir, error) {
	var err error
	if dirStruct.Hash == "" {
		_, err = dirStruct.setHash()
	} else {
		_, err = dirStruct.removeHash()
	}
	if err != nil {
		return dirStruct, err
	}
	dirStruct.setURL()
	return dirStruct, nil
}
