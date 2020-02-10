package files

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Checks if the given dest is a child directory
// of the base directory
func isSubPath(baseDir, dest string) bool {
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

// Checks and construct the dest
func ConstructPath(baseDir, dest string) (string, error) {
	if dest == "" {
		dest = baseDir
	}
	dest = filepath.Dir(dest)

	if !isSubPath(baseDir, dest) {
		return "", errors.New("unable to process the dest")
	}
	return dest + "/", nil
}

// Gets a list of directories and files
func GetDirectories(baseDir, dest string) ([]string, error) {
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

// Zips directory
func ZipDir(dir, dest, name string) (string, error) {
	destinationPath := dest + name + ".zip"
	if _, err := os.Stat(destinationPath); err == nil {
		return destinationPath, nil
	}
	err := RecursiveZip(dir, destinationPath)
	if err != nil {
		return "", err
	}
	return destinationPath, nil
}

//  Deletes the zip file with the provided name and destination path
func DeleteZip(dest, name string) error {
	destinationPath := dest + name + ".zip"
	err := os.Remove(destinationPath)
	return err
}

// Recursive zips the directory
func RecursiveZip(pathToZip, destinationPath string) error {
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	myZip := zip.NewWriter(destinationFile)
	err = filepath.Walk(
		pathToZip,
		func(filePath string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			relPath := strings.TrimPrefix(filePath, filepath.Dir(pathToZip))
			zipFile, err := myZip.Create(relPath)
			if err != nil {
				return err
			}
			fsFile, err := os.Open(filePath)
			if err != nil {
				return err
			}
			_, err = io.Copy(zipFile, fsFile)
			if err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return err
	}
	err = myZip.Close()
	if err != nil {
		return err
	}
	return nil
}
