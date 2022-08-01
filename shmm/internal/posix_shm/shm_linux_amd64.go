//go:build linux

package posix_shm

import (
	"os"
	"path/filepath"
)

// Create : inspired from github.com/fabiokung/shm
func Create(shmname string, size int64, perm os.FileMode) (*os.File, error) {
	name := filepath.Join("/dev/shm", shmname)
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, perm)
	if err != nil {
		return nil, err
	}

	err = file.Truncate(size)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Unlink : inspired from github.com/fabiokung/shm
func Unlink(shmname string) error {
	name := filepath.Join("/dev/shm", shmname)
	return os.Remove(name)
}

func setFirstCharToSlash(str string) (result string) {
	if str[0:1] != "/" {
		result = "/" + str
	} else {
		result = str
	}
	return
}
