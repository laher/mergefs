package mergefs

import (
	"errors"
	"github.com/rs/zerolog/log"
	"io/fs"
	"os"
	"sort"
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

// ReadDir reads from the directory, and produces a DirEntry array of different
// directories.
//
// It iterates through all different filesystems that exist in the mfs MergeFS
// filesystem slice and it identifies overlapping directories that exist in different
// filesystems
func (mfs MergedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	dirsMap := make(map[string]fs.DirEntry)
	for _, filesystem := range mfs.filesystems {
		if fsys, ok := filesystem.(fs.ReadDirFS); ok {
			dir, err := fsys.ReadDir(name)
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					log.Debug().Msgf("directory in filepath %s was not found in filesystem", name)
					continue
				}
				return nil, err
			}
			for _, v := range dir {
				if _, ok := dirsMap[v.Name()]; !ok {
					dirsMap[v.Name()] = v
				}
			}
			continue
		}

		file, err := filesystem.Open(name)
		if err != nil {
			log.Debug().Msgf("filepath %s was not found in filesystem", name)
			continue
		}

		dir, ok := file.(fs.ReadDirFile)
		if !ok {
			return nil, &fs.PathError{Op: "readdir", Path: name, Err: errors.New("not implemented")}
		}

		fsDirs, err := dir.ReadDir(-1)
		if err != nil {
			return nil, err
		}
		sort.Slice(fsDirs, func(i, j int) bool { return fsDirs[i].Name() < fsDirs[j].Name() })
		for _, v := range fsDirs {
			if _, ok := dirsMap[v.Name()]; !ok {
				dirsMap[v.Name()] = v
			}
		}
		if err := file.Close(); err != nil {
			log.Error().Msgf("failed to close file, err: %q", err)
		}
	}
	dirs := make([]fs.DirEntry, 0, len(dirsMap))

	for _, value := range dirsMap {
		dirs = append(dirs, value)
	}

	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
	return dirs, nil
}
