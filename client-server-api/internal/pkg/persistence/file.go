package persistence

import (
	"os"
)

type File struct {
	Path  string
	Flags int
}

func NewFile(path string, flags int) *File {
	return &File{
		Path:  path,
		Flags: flags,
	}
}

func (f File) Write(data []byte) error {
	fh, err := os.OpenFile(f.Path, f.Flags, 0644)

	if err != nil {
		return err
	}

	defer fh.Close()

	if _, err := fh.Write(data); err != nil {
		return err
	}

	return nil
}
