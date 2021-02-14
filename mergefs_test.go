package mergefs_test

import (
	"testing"
	"testing/fstest"

	"github.com/laher/mergefs"
)

func TestMergeFS(t *testing.T) {
	a := fstest.MapFS{"a": &fstest.MapFile{Data: []byte("text")}}
	b := fstest.MapFS{"b": &fstest.MapFile{Data: []byte("text")}}
	fs := mergefs.Merge(a, b)

	if _, err := fs.Open("a"); err != nil {
		t.Fatalf("file should exist")
	}
	if _, err := fs.Open("b"); err != nil {
		t.Fatalf("file should exist")
	}
}
