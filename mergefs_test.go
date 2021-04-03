package mergefs_test

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/laher/mergefs"
)

func TestMergeFS(t *testing.T) {
	a := fstest.MapFS{"a": &fstest.MapFile{Data: []byte("text")}}
	b := fstest.MapFS{"b/c": &fstest.MapFile{Data: []byte("text")}}
	mfs := mergefs.Merge(a, b)

	if _, err := mfs.Open("a"); err != nil {
		t.Fatalf("file should exist")
	}
	if _, err := mfs.Open("b/c"); err != nil {
		t.Fatalf("file should exist")
	}

	_, err := mfs.Open("b/d")
	if err == nil {
		t.Fatalf("file should not exist but nil error retuned")
	}
	if e, ok := err.(*fs.PathError); !ok {

		t.Fatalf("file should not exist: %T: %#v", err, err)
	} else {
		if e.Err != fs.ErrNotExist {
			t.Fatalf("error was not fs.ErrNotExist: %T", e.Err)
		}
	}
	fstest.TestFS(mfs, "a", "b/c")
}
