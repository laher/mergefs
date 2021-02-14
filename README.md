# Mergefs

A tiny go package which combines together fs.FS filesystems.

`mergefs.FS` looks through a slice of `fs.FS` filesytems in order to find a given file. It returns the first match, or `os.ErrNotExist`.

[![Go Reference](https://pkg.go.dev/badge/github.com/laher/chainfs.svg)](https://pkg.go.dev/github.com/laher/chainfs)
