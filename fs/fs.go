package fs

var fs = &osFileSystem{}

func WriteJSON(path string, data interface{}) error {
	return fs.WriteJSON(path, data)
}

func ReadJSON(path string, data interface{}) error {
	return fs.ReadJSON(path, data)
}

func Exists(path string) bool {
	return fs.Exists(path)
}

func MkdirP(path string) error {
	return fs.MkdirP(path)
}

func WriteFile(name string, writeFunc WriteFunc) error {
	return fs.WriteFile(name, writeFunc)
}

func ReadFile(name string, readFunc ReadFunc) error {
	return fs.ReadFile(name, readFunc)
}

func WriteFileNotExist(name string, writeFunc WriteFunc) error {
	return fs.WriteFileNotExist(name, writeFunc)
}
