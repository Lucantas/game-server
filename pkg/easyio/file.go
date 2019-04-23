package easyio

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FileName struct {
	Path string
}

// ReadFile takes a path and return the file content
// as slice of bytes using the ioutil's ReadAll method
// on a reader of type File, return the errors from the
// packages methods in case of any
func ReadFile(path string) ([]byte, error) {
	var content []byte
	file, err := os.Open(path)

	if err != nil {
		return content, err
	}

	//func (f *File) Read(b []byte) (n int, err error)
	content, err = ioutil.ReadAll(file)

	if err != nil {
		return content, err
	}

	return content, nil
}

func (f FileName) GetParent() string {
	var parent string
	// find the current directory
	// name can be achiveved with filepath.Base()
	base := path.Base(filepath.Dir(f.Path))
	// test if this name is repeated along the file path string

	if strings.Count(f.Path, base) == 1 {
		// the parent will be the first element of the result
		// of a split of the file path based on the current directory
		parent = strings.Split(f.Path, base)[0]
	} else if strings.Count(f.Path, base) > 1 {
		// the parent directory will be the second to last string
		// of the result of a split of the file path based on the current directory
		splitted := strings.SplitAfterN(f.Path, base, strings.Index(f.Path, base))
		splitted = append(splitted[:len(splitted)-1], splitted[len(splitted):]...)
		parent = strings.Join(splitted, "")
	}

	return parent + base
}
