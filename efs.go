package gojsonschema

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

type EmbeddedFS struct {
	fs *embed.FS
}

type EmbeddedFile struct {
	file fs.File
}

func (efs *EmbeddedFS) Open(name string) (http.File, error) {
	file, error := efs.fs.Open(name)
	if error != nil {
		return nil, error
	}
	return &EmbeddedFile{file: file}, nil
}

func (ef *EmbeddedFile) Close() error {
	return ef.file.Close()
}

func (ef *EmbeddedFile) Read(p []byte) (n int, err error) {
	n, err = ef.file.Read(p)
	return
}

func (ef *EmbeddedFile) Seek(offset int64, whence int) (int64, error) {
	return 0, fmt.Errorf("Seek is not supported by embedded file system")
}

func (ef *EmbeddedFile) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}

func (ef *EmbeddedFile) Stat() (fs.FileInfo, error) {
	return ef.file.Stat()
}
