package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/mitch292/janitor/internal/constants"
	"github.com/mitch292/janitor/internal/util"
)

// The InternalFileDirect or the "Janitor's closet" is where all the files will go.
// From there we symlink to the desired destination.
type internalFileDirectory struct {
	location string
}

// NewInternalFileDirectory will create the internal file directory, the "Janitor's closet".
// All files will be stored here, then symlinked to their ultimate destination.
func NewInternalFileDirectory() *internalFileDirectory {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	location := path.Join(home, constants.INTERNAL_DIR_NAME)

	if err := os.MkdirAll(location, os.FileMode(0700)); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return &internalFileDirectory{
		location: location,
	}
}

// Sync will take a file from the source and add it to the destination
func (f *internalFileDirectory) Sync(janitorFile *janitorFile) (err error) {
	if err = janitorFile.getFileDataFromSource(); err != nil {
		return
	}

	if err = janitorFile.writeFileDataToDestination(); err != nil {
		return
	}

	janitorFile.createSymlinkToDir(f)

	return
}

// Mop will remove broken symlinks from the internalFileDirectory
func (f *internalFileDirectory) Mop() (err error) {
	files, err := ioutil.ReadDir(f.location)
	if err != nil {
		return
	}

	for _, file := range files {
		symLinkPath := path.Join(f.location, file.Name())

		if err := util.RemoveSymlinkIfBroken(symLinkPath); err != nil {
			return err
		}
	}
	return
}

// Trash will remove all files managed by janitor.
func (f *internalFileDirectory) Trash() (err error) {
	files, err := ioutil.ReadDir(f.location)
	if err != nil {
		return
	}

	for _, file := range files {
		symLinkPath := path.Join(f.location, file.Name())

		util.RemoveSymlinkAndSourceFile(symLinkPath)
	}
	return
}
