package chainfs

import (
	"io/fs"
	"os"
)

type FS struct {
	Filesystems []fs.FS
}

// Open opens the named file.
func (cfs FS) Open(name string) (fs.File, error) {
	for _, fs := range cfs.Filesystems {
		file, err := fs.Open(name)
		if err == nil {
			return file, nil
		}
	}
	return nil, os.ErrNotExist
}
