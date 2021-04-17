package util

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestIsDirectoryWithDirectory(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatal("Could not create temporary testing directory.")
	}
	defer os.RemoveAll(tmpDir)

	isDir := IsDirectory(tmpDir)

	if !isDir {
		t.Fatalf("IsDirectory did not identify a directory properly: %s", tmpDir)
	}
}

func TestIsDirectoryWithFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal("Could not create a temporary testing file.")
	}
	defer os.Remove(tmpFile.Name())

	isDir := IsDirectory(tmpFile.Name())

	if isDir {
		t.Fatalf("IsDirectory identified a file as a directory: %s", tmpFile.Name())
	}
}

func TestFileExistsWithFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal("Could not create a temporary testing file.")
	}
	defer os.Remove(tmpFile.Name())

	fileExists := FileExists(tmpFile.Name())

	if !fileExists {
		t.Fatalf("FileExists did not identify a file that exists: %s", tmpFile.Name())
	}
}

func TestFileExistsWithNoFile(t *testing.T) {
	fileExists := FileExists("not-a-real-file")

	if fileExists {
		t.Fatal("FileExists identified a nonexistent file as existing")
	}
}

func TestAbsolutePathWithTilde(t *testing.T) {
	home, err := homedir.Dir()
	if err != nil {
		t.Fatal("Could not determine home directory while preparing test.")
	}

	relativePath := "~/testdir"
	absPath, err := AbsolutePath(relativePath)
	if err != nil {
		t.Fatalf("AbsolutePath threw an error when evaulating the path of %s", relativePath)
	}

	expectedPath := path.Join(home, "testdir")
	if absPath != expectedPath {
		t.Fatalf("AbsolutePath returned the wrong path for %s", relativePath)
	}
}

func TestAbsolutePathWithAbsolutePath(t *testing.T) {
	home, err := homedir.Dir()
	if err != nil {
		t.Fatal("Could not determine home directory while preparing test.")
	}

	tmpDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatal("Could not create temporary testing directory.")
	}
	defer os.RemoveAll(tmpDir)

	fullPath := path.Join(home, "tmpDir")

	absPath, err := AbsolutePath(fullPath)
	if err != nil {
		t.Fatalf("AbsolutePath threw an error when evaulating the path of %s", fullPath)
	}

	if absPath != fullPath {
		t.Fatalf("AbsolutePath returned the wrong path for %s", fullPath)
	}
}

func TestRemoveSymlinkIfBrokenWithBrokenSymlink(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal("Could not create a temporary testing file.")
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal("Could not get the cwd while preparing test.")
	}

	symlinkPath := path.Join(cwd, "link")

	os.Symlink(tmpFile.Name(), symlinkPath)
	defer os.Remove(symlinkPath)

	os.Remove(tmpFile.Name()) // breaks symlink
	RemoveSymlinkIfBroken(symlinkPath)

	if FileExists(symlinkPath) {
		t.Fatalf("RemoveSylinkIfBroken did not removed a symlink (%s) when the source file was removed.", symlinkPath)
	}
}

func TestRemoveSymlinkIfBrokenWithValidSymlink(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal("Could not create a temporary testing file.")
	}
	defer os.Remove(tmpFile.Name())

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal("Could not get the cwd while preparing test.")
	}

	symlinkPath := path.Join(cwd, "link")

	os.Symlink(tmpFile.Name(), symlinkPath)
	defer os.Remove(symlinkPath)

	RemoveSymlinkIfBroken(symlinkPath)

	if !FileExists(symlinkPath) {
		t.Fatalf("RemoveSylinkIfBroken removed a symlink (%s) when the source file (%s) still existed.", symlinkPath, tmpFile.Name())
	}
}

func TestRemoveSymlinkAndSourceFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal("Could not create a temporary testing file.")
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal("Could not get the cwd while preparing test.")
	}

	symlinkPath := path.Join(cwd, "link")
	os.Symlink(tmpFile.Name(), symlinkPath)

	RemoveSymlinkAndSourceFile(symlinkPath)

	if FileExists(symlinkPath) {
		t.Fatalf("RemoveSymlinkAndSourceFile failed to remove the symlink %s", symlinkPath)
	}

	if FileExists(tmpFile.Name()) {
		t.Fatalf("RemoveSymlinkAndSourceFile failed to remove the file %s", tmpFile.Name())
	}
}
