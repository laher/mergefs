# Mergefs

A tiny go package which combines together fs.FS filesystems.

`mergefs.FS` looks through a slice of `fs.FS` filesytems in order to find a given file. It returns the first match, or `os.ErrNotExist`.

[![Go Reference](https://pkg.go.dev/badge/github.com/laher/mergefs.svg)](https://pkg.go.dev/github.com/laher/mergefs)

# Related work

mergefs could be used to overlay multiple fs.FS filesystems on top of each other.

 * [marshalfs](https://pkg.go.dev/github.com/laher/marshalfs) for backing a fileystem with your objects.
 * Standard Library:
   * [dirfs](https://tip.golang.org/pkg/os/) contains `os.DirFS` - this 'default' implementation is backed by an actual filesystem.
   * [testfs](https://tip.golang.org/pkg/testing/fstest/) contains a memory-map implementation and a testing tool. The standard library contains a few other fs.FS implementations (like 'zip')
   * [embedfs](https://tip.golang.org/pkg/embed/) provides access to files embedded in the running Go program.
 * An earlier work, [afero](https://github.com/spf13/afero) is a filesystem abstraction for Go, which has been the standard for filesystem abstractions up until go1.15. It's read-write, and it's a mature project. The interfaces look very different (big with lots of methods), so it's not really compatible.
 * [s3fs](https://github.com/jszwec/s3fs) is a fs.FS backed by the AWS S3 client
 * [mergefs](https://github.com/laher/mergefs) merge `fs.FS` filesystems together so that your FS can easily read from multiple sources.
 * [hashfs](https://pkg.go.dev/github.com/benbjohnson/hashfs) appends SHA256 hashes to filenames to allow for aggressive HTTP caching.
