package mergefs_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/laher/mergefs"
	"github.com/stretchr/testify/require"
)

func TestMergeFS(t *testing.T) {

	t.Run("testing different filesystems", func(t *testing.T) {
		a := fstest.MapFS{"a": &fstest.MapFile{Data: []byte("text")}}
		b := fstest.MapFS{"b": &fstest.MapFile{Data: []byte("text")}}
		filesystem := mergefs.Merge(a, b)

		if _, err := filesystem.Open("a"); err != nil {
			t.Fatalf("file should exist")
		}
		if _, err := filesystem.Open("b"); err != nil {
			t.Fatalf("file should exist")
		}
	})

	var filePaths = []struct {
		path           string
		dirArrayLength int
		child          string
	}{
		// MapFS takes in account the current directory in addition to all included directories and produces a "" dir
		{"a", 1, "z"},
		{"a/z", 1, "bar.cue"},
		{"b", 1, "z"},
		{"b/z", 1, "foo.cue"},
	}

	tempDir := os.DirFS(filepath.Join("testdata"))
	a := fstest.MapFS{
		"a":           &fstest.MapFile{Mode: fs.ModeDir},
		"a/z":         &fstest.MapFile{Mode: fs.ModeDir},
		"a/z/bar.cue": &fstest.MapFile{Data: []byte("bar")},
	}

	filesystem := mergefs.Merge(tempDir, a)

	t.Run("testing mergefs.ReadDir", func(t *testing.T) {
		for _, fp := range filePaths {
			t.Run("testing path: "+fp.path, func(t *testing.T) {
				dirs, err := fs.ReadDir(filesystem, fp.path)
				require.NoError(t, err)
				require.Len(t, dirs, fp.dirArrayLength)

				for i := 0; i < len(dirs); i++ {
					require.Equal(t, dirs[i].Name(), fp.child)
				}
			})
		}
	})

	t.Run("testing mergefs.Open", func(t *testing.T) {
		data := make([]byte, 3)
		file, err := filesystem.Open("a/z/bar.cue")
		require.NoError(t, err)

		_, err = file.Read(data)
		require.NoError(t, err)
		require.Equal(t, "bar", string(data))

		file, err = filesystem.Open("b/z/foo.cue")
		require.NoError(t, err)

		_, err = file.Read(data)
		require.NoError(t, err)
		require.Equal(t, "foo", string(data))

		err = file.Close()
		require.NoError(t, err)
	})
}
