package fs

import (
	"encoding/json"
	"io"
	"os"
)

type (
	WriteFunc func(writer io.Writer) error
	ReadFunc func(reader io.Reader) error
)

// FileSystem ...
type FileSystem interface {
	Exists(path string) bool
	MkdirP(path string) error
	ReadJSON(path string, data interface{}) error
	WriteJSON(path string, data interface{}) error
	WriteFileNotExist(name string, writeFunc WriteFunc) error
	ReadFile(name string, readFunc ReadFunc) error
	WriteFile(name string, writeFunc WriteFunc) error
}

// File ...
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	Stat() (os.FileInfo, error)
}

type osFileSystem struct{}

func (o *osFileSystem) ReadJSON(path string, data interface{}) error {
	return o.ReadFile(path, func(reader io.Reader) error {
		err := json.NewDecoder(reader).Decode(data)
		if err != nil {
			return err
		}
		return nil
	})
}

func (o *osFileSystem) WriteJSON(path string, data interface{}) error {
	return o.WriteFile(path, func(writer io.Writer) error {
		err := json.NewEncoder(writer).Encode(data)
		if err != nil {
			return err
		}
		return nil
	})
}

func (*osFileSystem) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (*osFileSystem) MkdirP(path string) error {
	if fileInfo, err := os.Stat(path); os.IsNotExist(err) || (fileInfo != nil && !fileInfo.IsDir()) {
		return os.MkdirAll(path, 0751)
	}
	return nil
}

func (*osFileSystem) ReadFile(name string, readFunc ReadFunc) error {
	file, err := os.OpenFile(name, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()
	return readFunc(file)
}

func (*osFileSystem) WriteFile(name string, writeFunc WriteFunc) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return writeFunc(file)
}

func (o *osFileSystem) WriteFileNotExist(name string, writeFunc WriteFunc) error {
	_, err := os.Stat(name)
	if err != nil {
		return o.WriteFile(name, writeFunc)
	}
	return nil
}
