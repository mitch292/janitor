package files

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/mitch292/janitor/internal/util"

	homedir "github.com/mitchellh/go-homedir"
)

const INTERNAL_DIR_NAME = ".janitors_closet"
const INTERNAL_ERROR_LOG_FILE = "_janitors_error_log"

// The InternalFileDirect or the "Janitor's closet" is where all the files will go.
// From there we symlink to the desired destination.
type InternalFileDirectory struct {
	location         string
	errorLogLocation string
}

// CreateInternalFileDirectory will create the internal file directory, the "Janitor's closet".
// All files will be stored here, then symlinked to their ultimate destination.
func CreateInternalFileDirectory() *InternalFileDirectory {
	f := new(InternalFileDirectory)

	// create the internal file directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	f.location = path.Join(home, INTERNAL_DIR_NAME)

	if err := os.MkdirAll(f.location, os.FileMode(0777)); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	f.errorLogLocation = INTERNAL_ERROR_LOG_FILE

	return f
}

// Sync will take a file from the source and add it to the destination
func (fileDir *InternalFileDirectory) Sync(name, source, destination string) (err error) {
	// TODO: Should the interaction with an http client be somewhere else?
	resp, err := http.Get(source)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	destinationIsDirectory := util.IsDirectory(destination)

	var ultimateFileDestination string
	if destinationIsDirectory {
		ultimateFileDestination = path.Join(destination, name)
	} else {
		ultimateFileDestination = destination
	}

	out, err := os.Create(ultimateFileDestination)
	if err != nil {
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return
	}

	internalFileLoc := path.Join(fileDir.location, name)
	os.Symlink(ultimateFileDestination, internalFileLoc)

	return
}

func (fileDir *InternalFileDirectory) WriteError(message string) error {
	// If the file doesn't exist, create it, or append to the file
	errorLogFile, err := os.OpenFile(path.Join(fileDir.location, fileDir.errorLogLocation), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer errorLogFile.Close()

	logger := log.New(errorLogFile, "janitor: ", log.LstdFlags)
	logger.Println(message)

	return nil
}
