package file

import (
	"errors"
	"io"
	"os"
)

func CheckPath(path string) error {
	empty, err := IsEmpty(path)
	if os.IsNotExist(err) {
		return ExistOrCreate(path)
	}

	if err != nil {
		return err
	}

	if !empty {
		return errors.New("path is not empty")
	}

	return nil
}

func ExistOrCreate(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// attempt creation
		err := os.MkdirAll(path, 0700)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func ListDirectories(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs = make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs, nil

}