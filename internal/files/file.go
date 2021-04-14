package files

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/mitch292/janitor/internal/util"
)

type janitorFile struct {
	sourceLocation      string
	destinationLocation string
	name                string
	contents            []byte
}

// NewJanitorFile will create a new janitorFile struct. This will hep us  interact
// with the source file and destination
func NewJanitorFile(name, sourceLocation, destinationLocation string) (*janitorFile, error) {
	destination, err := util.AbsolutePath(destinationLocation)
	if err != nil {
		return &janitorFile{}, err
	}
	return &janitorFile{
		name:                name,
		sourceLocation:      sourceLocation,
		destinationLocation: destination,
		contents:            make([]byte, 0),
	}, nil
}

func (f *janitorFile) getFileDataFromSource() (err error) {
	resp, err := http.Get(f.sourceLocation)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	f.contents, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

func (f *janitorFile) writeFileDataToDestination() (err error) {
	out, err := os.Create(f.destinationLocation)
	if err != nil {
		return
	}
	defer out.Close()

	if _, err = out.Write(f.contents); err != nil {
		return
	}

	if err = out.Sync(); err != nil {
		return
	}

	return
}

func (f *janitorFile) createSymlinkToDir(directory *internalFileDirectory) {
	os.Symlink(f.destinationLocation, path.Join(directory.location, f.name))
}
