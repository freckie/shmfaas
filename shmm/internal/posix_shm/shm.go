package shm

import (
	"os"
	"syscall"
	"unsafe"
)

// Create : inspired from github.com/fabiokung/shm
func Create(shmname string, size int64, perm os.FileMode) (*os.File, error) {
	name, err := syscall.BytePtrFromString(shmname) // char *
	if err != nil {
		return nil, err
	}

	file, _, errNo := syscall.Syscall(syscall.SYS_SHM_OPEN,
		uintptr(unsafe.Pointer(name)),
		uintptr(os.O_RDWR|os.O_CREATE),
		uintptr(perm),
	)
	if errNo != 0 {
		return nil, errNo
	}

	err = syscall.Ftruncate(int(file), size)
	if err != nil {
		return nil, err
	}

	return os.NewFile(file, shmname), nil
}

// Unlink : inspired from github.com/fabiokung/shm
func Unlink(shmname string) error {
	name, err := syscall.BytePtrFromString(shmname)
	if err != nil {
		return err
	}

	_, _, errNo := syscall.Syscall(syscall.SYS_SHM_UNLINK,
		uintptr(unsafe.Pointer(name)),
		uintptr(0),
		uintptr(0),
	)
	if errNo != 0 {
		return errNo
	}

	return nil
}
