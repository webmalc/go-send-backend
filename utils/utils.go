package utils

import (
	"archive/zip"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ProcessFatalError checks  fatal errors
func ProcessFatalError(err error) {
	if err != nil {
		panic(fmt.Errorf("error: %s", err))
	}
}

// GeneateUUID generates UUID
func GeneateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// GenerateRandomString generates a random string
func GenerateRandomString(lenght int) string {
	result := ""
	for i := 0; i < lenght; i++ {
		result += GeneateUUID()
	}
	return result
}

// ZipDir zips directory
func ZipDir(dir string, dest string, name string) (string, error) {
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

// DeleteFile deletes file
func DeleteFile(dest string, name string) error {
	destinationPath := dest + name + ".zip"
	err := os.Remove(destinationPath)
	return err
}

// RecursiveZip recursive zips the directory
func RecursiveZip(pathToZip, destinationPath string) error {
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	myZip := zip.NewWriter(destinationFile)
	err = filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, err error) error {
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
