package errs

import (
	"fmt"
	"runtime"
)

type AssertError struct {
	checkerError
}

func (ae *AssertError) Error() string {
	if ae.file == "" {
		return fmt.Sprintf("assertion failed at <unknown>")
	} else {
		return fmt.Sprintf("assertion failed at %s:%d", ae.file, ae.line)
	}
}

func Assert(cond bool, args ...interface{}) {
	if cond {
		return
	}
	ae := AssertError{checkerError: *newCheckerError(1, nil, nil, args)}
	ae.stackTrace = make([]byte, 1<<20)
	runtime.Stack(ae.stackTrace, false)
	panic(&ae)
}
