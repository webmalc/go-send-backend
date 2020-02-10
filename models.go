package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"
	"github.com/webmalc/go-send-backend/config"
	"github.com/webmalc/go-send-backend/files"
	"github.com/webmalc/go-send-backend/utils"
)

const hashLength = 5

// The directory struct
type Dir struct {
	Path   string         `json:"path" binding:"required"`
	Hash   string         `json:"hash" binding:"required"`
	URL    string         `json:"url" binding:"required"`
	Db     *redis.Client  `json:"-"`
	Logger *log.Logger    `json:"-"`
	Config *config.Config `json:"-"`
}

// Constructs a Dir structure
func (dir *Dir) constructor(
	dirPath string,
	db *redis.Client,
	logger *log.Logger,
	conf *config.Config,
) {
	dir.Path = dirPath
	dir.Db = db
	dir.Logger = logger
	dir.Config = conf
	dir.setHashFromDB()
	dir.setURL()
}

// Sets the Dir hash from the database
func (dir *Dir) setHashFromDB() {
	hash, _ := dir.Db.Get(dir.Path).Result()
	dir.Hash = hash
}

// Sets URL for the Dir
func (dir *Dir) setURL() {
	dir.URL = ""
	if dir.Hash != "" && dir.Path != "" {
		pattern := "%s/public/get/%s/%s"
		host := dir.Config.Host
		base := base64.StdEncoding.EncodeToString([]byte(dir.Path))
		dir.URL = fmt.Sprintf(pattern, host, dir.Hash, base)
	}
}

// Removes the Dir hash
func (dir *Dir) removeHash() error {
	_, err := dir.Db.Del(dir.Path).Result()
	if err != nil {
		return err
	}
	dir.Logger.Printf(
		"[INFO] Deleting the zip %s%s.zip",
		dir.Config.ZipPath, dir.Hash)
	err = files.DeleteZip(dir.Config.ZipPath, dir.Hash)
	if err != nil {
		dir.Logger.Printf(
			"[ERROR] Failed to delete the zip %s%s.zip",
			dir.Config.ZipPath, dir.Hash)
	}
	dir.Hash = ""
	return nil
}

// Sets the Dir hash
func (dir *Dir) setHash() error {
	hash := utils.GenerateRandomString(hashLength)
	err := dir.Db.Set(dir.Path, hash, 0).Err()
	if err != nil {
		return err
	}
	dir.Hash = hash
	_, err = dir.generateZip()
	dir.Logger.Printf(
		"[INFO] Generating a zip for the directory %s in %s%s.zip",
		dir.Path, dir.Config.ZipPath, dir.Hash)
	if err != nil {
		dir.Logger.Printf(
			"[ERROR] Failed to generate a zip for the directory %s in %s%s.zip",
			dir.Path, dir.Config.ZipPath, dir.Hash)
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

// Generates a zip archive for the Dir object
func (dir *Dir) generateZip() (string, error) {
	if dir.Path != "" && dir.Hash != "" {
		return files.ZipDir(dir.Path, dir.Config.ZipPath, dir.Hash)
	}
	return "", errors.New("unable to generate a zip archive")
}
