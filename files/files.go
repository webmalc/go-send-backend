package files

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// isSubPath checks if the given path is a child directory
// of the base directory
func isSubPath(baseDir string, path string) bool {
	path, _ = filepath.EvalSymlinks(path)
	rel, err := filepath.Rel(baseDir, path)
	if err != nil {
		return false
	}

	if strings.Contains(rel, "..") {
		return false
	}
	return true
}

// ConstructPath checks and construct the path
func ConstructPath(baseDir string, path string) (string, error) {
	if path == "" {
		path = baseDir
	}
	path = filepath.Dir(path)

	if !isSubPath(baseDir, path) {
		return "", errors.New("unable to process the path")
	}
	return path, nil
}

// GetDirectories gets a list of directories and files
func GetDirectories(baseDir string, path string) ([]string, error) {
	path, err := ConstructPath(baseDir, path)
	var result []string
	if err != nil {
		return result, err
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return result, errors.New("unable to read the directory")
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			result = append(result, fileInfo.Name())
		}
	}
	return result, nil
}
