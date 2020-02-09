package files

import (
	"os"
	"path/filepath"
	"testing"
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
	if !isSubPath(etc, kernel) {
		t.Errorf("%s contains %s", kernel, etc)
	}
	if isSubPath(etc, invalid) {
		t.Errorf("%s does not contain %s", invalid, etc)
	}
}

// Should construct a path if the given path contains the base directory
func TestConstructPath(t *testing.T) {
	r, err := ConstructPath(etc, kernel)

	if err != nil || r != "/etc/kernel/" {
		t.Errorf("%s contains %s, got: %s", kernel, etc, r)
	}
	if _, err := ConstructPath(etc, invalid); err == nil {
		t.Errorf("%s does not contains %s", invalid, etc)
	}
}

// Should return a slice of directories for the given destination path
func TestGetDirectories(t *testing.T) {
	dirs, err := GetDirectories(baseDir, "")
	if err != nil {
		t.Errorf("error %s", err)
	}
	set := make(map[string]bool)
	for _, v := range dirs {
		set[v] = true
	}
	if !set[baseDir+"config/"] {
		t.Errorf("unable to find the config directory. got %s", dirs)
	}
	if !set[baseDir+"files/"] {
		t.Errorf("unable to find the files directory. got %s", dirs)
	}
	if !set[baseDir+"utils/"] {
		t.Errorf("unable to find the files directory. got %s", dirs)
	}

}

// Should return errors when the provided path is incorrect
func TestGetDirectoriesError(t *testing.T) {
	dirs, err := GetDirectories(baseDir, etc)
	if err == nil {
		t.Errorf("invalid subpath has been provided. got %s", dirs)
	}
	dirs, err = GetDirectories("/root/", "")
	if err == nil {
		t.Errorf("nonexistent path has been provided. got %s", dirs)
	}
}

// Should zip the provided directory
func TestZipDirAndDelete(t *testing.T) {
	zip, err := ZipDir(baseDir+"utils/", baseDir, "test")
	if err != nil {
		t.Errorf("error %s", err)
	}
	if zip != baseDir+"test.zip" {
		t.Errorf("invalid zip path %s", zip)
	}
	zip, _ = ZipDir(baseDir+"utils/", baseDir, "test")
	if zip != baseDir+"test.zip" {
		t.Errorf("invalid zip path %s", zip)
	}
	_ = DeleteZip(baseDir, "test")
	if _, err := os.Stat(zip); err == nil {
		t.Errorf("the zip file has not been deleted. file: %s", zip)
	}
}

// Should return errors when the provided path is incorrect
func TestZipDirErrors(t *testing.T) {
	zip, err := ZipDir(baseDir, baseDir+"zip/", "test")
	if err == nil {
		t.Errorf("an error should have been returned. file: %s", zip)
	}
}

// Should return errors when the provided path is incorrect
func TestRecursiveZipErrors(t *testing.T) {
}
