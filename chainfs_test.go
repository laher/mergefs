package chainfs

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestChainFS(t *testing.T) {
	a := fstest.MapFS{"a": &fstest.MapFile{Data: []byte("text")}}
	b := fstest.MapFS{"b": &fstest.MapFile{Data: []byte("text")}}
	fs := FS{Filesystems: []fs.FS{a, b}}

	if _, err := fs.Open("a"); err != nil {
		t.Fatalf("file should exist")
	}
	if _, err := fs.Open("b"); err != nil {
		t.Fatalf("file should exist")
	}
}
