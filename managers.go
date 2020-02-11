package main

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/go-redis/redis/v7"
	"github.com/webmalc/go-send-backend/config"
)

type DirManager struct {
	Db     *redis.Client
	Logger *log.Logger
	Config *config.Config
}

// Returns the controller object
func NewManager(
	db *redis.Client,
	logger *log.Logger,
	conf *config.Config,
) *DirManager {
	return &DirManager{
		Db:     db,
		Logger: logger,
		Config: conf,
	}
}

// Returns a new constructed Dir object
func (manager *DirManager) getDir(dir string) Dir {
	dirStruct := Dir{}
	dirStruct.constructor(dir, manager.Db, manager.Logger, manager.Config)
	return dirStruct
}

// Constructs the Dir slice from paths
func (manager *DirManager) constructDirsSlice(dirs []string) []Dir {
	var results []Dir
	for _, dir := range dirs {
		results = append(results, manager.getDir(dir))
	}
	return results
}

// Toggles hash for the directory
func (manager *DirManager) toggleDirHash(dir string) (Dir, error) {
	dirStruct := manager.getDir(dir)
	err := dirStruct.toggleHash()
	if err != nil {
		return dirStruct, err
	}
	return dirStruct, nil
}

// Gets a Dir structure by the hash
func (manager *DirManager) GetDirByHash(hash, base string) (Dir, error) {
	decoded, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return Dir{}, err
	}
	key := string(decoded)
	dir := manager.getDir(key)
	dbHash, err := manager.Db.Get(key).Result()
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
