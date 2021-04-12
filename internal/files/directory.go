package files

import (
	"fmt"
	"log"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

const INTERNAL_DIR_NAME = ".janitors_closet"
const INTERNAL_ERROR_LOG_FILE = ".janitors_error_log"

// The InternalFileDirect or the "Janitor's closet" is where all the files will go.
// From there we symlink to the desired destination.
type internalFileDirectory struct {
	location         string
	errorLogLocation string
}

// CreateInternalFileDirectory will create the internal file directory, the "Janitor's closet".
// All files will be stored here, then symlinked to their ultimate destination.
func NewInternalFileDirectory() *internalFileDirectory {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	location := path.Join(home, INTERNAL_DIR_NAME)

	if err := os.MkdirAll(location, os.FileMode(0700)); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return &internalFileDirectory{
		location:         location,
		errorLogLocation: path.Join(location, INTERNAL_ERROR_LOG_FILE),
	}
}

// Sync will take a file from the source and add it to the destination
func (f *internalFileDirectory) Sync(janitorFile *janitorFile) (err error) {
	if err = janitorFile.GetFileDataFromSource(); err != nil {
		return
	}

	if err = janitorFile.WriteFileDataToDestination(); err != nil {
		return
	}

	janitorFile.CreateSymlinkToDir(f)

	return
}

// WriteErrorToLog will write an error to the internal janitor's closet log (.janitors_closet/.janitors_error_log)
func (f *internalFileDirectory) WriteErrorToLog(message string) (err error) {
	errorLogFile, err := os.OpenFile(f.errorLogLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer errorLogFile.Close()

	logger := log.New(errorLogFile, "janitor: ", log.LstdFlags)
	logger.Println(message)

	return
}
