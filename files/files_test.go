package files

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dir string
var baseDir string
var etc string
var kernel string
var invalid string

// Initializes the main variables
func init() {
	dir, _ = os.Getwd()
	baseDir = filepath.Dir(dir) + "/"
	etc = "/etc/"
	kernel = "/etc/kernel/"
	invalid = "/foo/bar/"
}

// Should return true if the given path contains the base directory
func TestIsSubPath(t *testing.T) {
	assert.True(t, isSubPath(etc, kernel))
	assert.False(t, isSubPath(etc, invalid))
}

// Should construct a path if the given path contains the base directory
func TestConstructPath(t *testing.T) {
	r, err := ConstructPath(etc, kernel)

	assert.Nil(t, err)
	assert.Equal(t, r, "/etc/kernel/")

	_, err = ConstructPath(etc, invalid)
	assert.NotNil(t, err)

	_, err = ConstructPath(etc, "kernel")
	assert.Nil(t, err)
	assert.Equal(t, r, "/etc/kernel/")
}

// Should return a slice of directories for the given destination path
func TestGetDirectories(t *testing.T) {
	dirs, err := GetDirectories(baseDir, "")

	assertT := assert.New(t)
	assertT.Nil(err)
	assertT.Contains(dirs, baseDir+"config/")
	assertT.Contains(dirs, baseDir+"files/")
	assertT.Contains(dirs, baseDir+"utils/")
}

// Should return errors when the provided path is incorrect
func TestGetDirectoriesError(t *testing.T) {
	_, err := GetDirectories(baseDir, etc)
	assert.NotNil(t, err)
	_, err = GetDirectories("/root/", "")
	assert.NotNil(t, err)
}

// Should zip the provided directory
func TestZipDirAndDelete(t *testing.T) {
	zip, err := ZipDir(baseDir+"utils/", baseDir, "test")

	assertT := assert.New(t)
	assertT.Nil(err)
	assertT.Equal(zip, baseDir+"test.zip")

	zip, _ = ZipDir(baseDir+"utils/", baseDir, "test")

	assertT.Equal(zip, baseDir+"test.zip")
	_ = DeleteZip(baseDir, "test")

	_, err = os.Stat(zip)
	assertT.NotNil(err)
}

// Should return errors when the provided path is incorrect
func TestZipDirErrors(t *testing.T) {
	_, err := ZipDir(baseDir, baseDir+"zip/", "test")
	assert.NotNil(t, err)
}
