package main

import (
	"encoding/base64"
	"fmt"

	"github.com/webmalc/go-send-backend/files"
	"github.com/webmalc/go-send-backend/utils"
)

const hashLength = 5

// The directory struct
type Dir struct {
	Path string
	Hash string
	URL  string
}

// Constructs a Dir structure
func (dir *Dir) constructor(dirPath string) {
	dir.Path = dirPath
	dir.setHashFromDB()
	dir.setURL()
}

// Sets the Dir hash from the database
func (dir *Dir) setHashFromDB() {
	hash, _ := db.Get(dir.Path).Result()
	dir.Hash = hash
}

// Sets URL for the Dir
func (dir *Dir) setURL() {
	dir.URL = ""
	if dir.Hash != "" && dir.Path != "" {
		pattern := "%s/public/get/%s/%s"
		host := configuration.Host
		base := base64.StdEncoding.EncodeToString([]byte(dir.Path))
		dir.URL = fmt.Sprintf(pattern, host, dir.Hash, base)
	}
}

// Removes the Dir hash
func (dir *Dir) removeHash() error {
	_, err := db.Del(dir.Path).Result()
	if err != nil {
		return err
	}
	_ = files.DeleteZip(configuration.ZipPath, dir.Hash)
	dir.Hash = ""
	return nil
}

// Sets the Dir hash
func (dir *Dir) setHash() error {
	hash := utils.GenerateRandomString(hashLength)
	err := db.Set(dir.Path, hash, 0).Err()
	if err != nil {
		return err
	}
	dir.Hash = hash
	_, err = generateZip(dir)
	if err != nil {
		return err
	}
	return nil
}

// Toggles the Dir hash
func (dir *Dir) toggleHash() error {
	var err error
	if dir.Hash == "" {
		err = dir.setHash()
	} else {
		err = dir.removeHash()
	}
	if err != nil {
		return err
	}
	dir.setURL()
	return nil
}
