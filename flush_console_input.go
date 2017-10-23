package flushconin

import (
	"github.com/jeet-parekh/winapi"
)

const _SUCCESS string = "The operation completed successfully."

// FlushConsoleInput flushes the console input buffer. It returns an error if unsuccessful.
func FlushConsoleInput() (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = panicErr.(error)
		}
	}()

	MustFlushConsoleInput()
	return
}

// MustFlushConsoleInput flushes the console input buffer. It panics if unsuccessful.
func MustFlushConsoleInput() {
	handle, err := winapi.GetStdHandle(winapi.STD_INPUT_HANDLE)
	if err.Error() != _SUCCESS {
		panic(err)
	}
	_, err = winapi.FlushConsoleInputBuffer(handle)
	if err.Error() != _SUCCESS {
		panic(err)
	}
}
