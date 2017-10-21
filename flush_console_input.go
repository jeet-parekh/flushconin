package flushconin

import (
	"github.com/jeet-parekh/winapi"
)

const _SUCCESS string = "The operation completed successfully."

// FlushConsoleInput flushes the console input buffer. It returns an error if unsuccessful.
func FlushConsoleInput() error {
	handle, err := winapi.GetStdHandle(winapi.STD_INPUT_HANDLE)
	if err.Error() != _SUCCESS {
		return err
	}
	_, err = winapi.FlushConsoleInputBuffer(handle)
	if err.Error() != _SUCCESS {
		return err
	}
	return nil
}

// MustFlushConsoleInput flushes the console input buffer. It panics if unsuccessful.
func MustFlushConsoleInput() {
	handle, err := winapi.GetStdHandle(winapi.STD_INPUT_HANDLE)
	if err.Error() != _SUCCESS {
		panic(err.Error())
	}
	_, err = winapi.FlushConsoleInputBuffer(handle)
	if err.Error() != _SUCCESS {
		panic(err.Error())
	}
}
