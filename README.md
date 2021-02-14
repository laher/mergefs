# Mergefs

A tiny go package which combines together fs.FS filesystems.

`mergefs.FS` looks through a slice of `fs.FS` filesytems in order to find a given file. It returns the first match, or `os.ErrNotExist`.

[![Go Reference](https://pkg.go.dev/badge/github.com/laher/mergefs.svg)](https://pkg.go.dev/github.com/laher/mergefs)


# Related work

This functionality is also bundled into [marshalfs](https://pkg.go.dev/github.com/laher/marshalfs).
 * [ ](
 * [testfs](https://tip.golang.org/pkg/testing/fstest/) contains a memory-map implementation
 * [s3fs](https://github.com/jszwec/s3fs) is a fs.FS backed by the AWS S3 client
