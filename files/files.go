package files

import (
	"errors"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

// isSubPath checks if the given dest is a child directory
// of the base directory
func isSubPath(baseDir string, dest string) bool {
	dest, _ = filepath.EvalSymlinks(dest)
	rel, err := filepath.Rel(baseDir, dest)
	if err != nil {
		return false
	}

	if strings.Contains(rel, "..") {
		return false
	}
	return true
}

// ConstructPath checks and construct the dest
func ConstructPath(baseDir string, dest string) (string, error) {
	if dest == "" {
		dest = baseDir
	}
	dest = filepath.Dir(dest)

	if !isSubPath(baseDir, dest) {
		return "", errors.New("unable to process the dest")
	}
	return dest + "/", nil
}

// GetDirectories gets a list of directories and files
func GetDirectories(baseDir string, dest string) ([]string, error) {
	dest, err := ConstructPath(baseDir, dest)
	var result []string
	if err != nil {
		return result, err
	}
	files, err := ioutil.ReadDir(dest)
	if err != nil {
		return result, errors.New("unable to read the directory")
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			dirPath := path.Join(dest, fileInfo.Name()) + "/"
			result = append(result, dirPath)
		}
	}
	return result, nil
}
