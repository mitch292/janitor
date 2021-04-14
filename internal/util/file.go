package util

import (
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// IsDirectory determines if a given file path is a directory or not.
func IsDirectory(path string) bool {
	// maybe a helper for all this logic in util?
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	mode := fileInfo.Mode()

	return mode.IsDir()
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// AbsolutePath will take a path string and convert it to an absolute path.
// If an absolute path is given, it will return the same path.
func AbsolutePath(filename string) (string, error) {
	if strings.HasPrefix(filename, "~") {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		return strings.Replace(filename, "~", home, 1), nil
	}
	return filename, nil
}
