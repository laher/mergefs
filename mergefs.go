package mergefs

import (
	"io/fs"
	"os"
)

// Merge filesystems
func Merge(filesystems ...fs.FS) fs.FS {
	return MergedFS{filesystems: filesystems}
}

// MergedFS combines filesystems. Each filesystem can serve different paths. The first FS takes precedence
type MergedFS struct {
	filesystems []fs.FS
}

// Open opens the named file.
func (mfs MergedFS) Open(name string) (fs.File, error) {
	for _, fs := range mfs.filesystems {
		file, err := fs.Open(name)
		if err == nil { // TODO should we return early when it's not an os.ErrNotExist? Should we offer options to decide this behaviour?
			return file, nil
		}
	}
	return nil, os.ErrNotExist
}
