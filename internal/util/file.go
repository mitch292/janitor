package util

import (
	"os"
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
