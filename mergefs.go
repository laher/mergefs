package mergefs

import (
	"io/fs"
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
	for _, mfs := range mfs.filesystems {
		file, err := mfs.Open(name)
		if err == nil { // TODO should we return early when it's not an os.ErrNotExist? Should we offer options to decide this behaviour?
			return file, nil
		}
		if e, ok := err.(*fs.PathError); ok {
			if e.Err != fs.ErrNotExist {
				return nil, err
			}
		}
	}
	return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
}
