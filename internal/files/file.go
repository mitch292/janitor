package files

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type janitorFile struct {
	sourceLocation      string
	destinationLocation string
	name                string
	contents            []byte
}

// NewJanitorFile will create a new janitorFile struct.
func NewJanitorFile(name, sourceLocation, destinationLocation string) *janitorFile {
	return &janitorFile{
		name:                name,
		sourceLocation:      sourceLocation,
		destinationLocation: destinationLocation,
		contents:            make([]byte, 0),
	}
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
