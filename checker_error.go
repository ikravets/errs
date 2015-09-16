package errs

import (
	"fmt"
	"runtime"
	"strconv"
)

type checkerError struct {
	err        error
	args       []interface{}
	stackTrace []byte
	file       string
	line       int
	checker    Checker
}

func newCheckerError(callerDepth int, checker Checker, err error, args []interface{}) *checkerError {
	e := &checkerError{
		err:     err,
		args:    args,
		checker: checker,
	}
	_, e.file, e.line, _ = runtime.Caller(callerDepth + 1)
	return e
}
func (e *checkerError) Error() string {
	fileStr, lineStr, errStr := "<?>", "<?>", "<nil>"
	if e.file != "" {
		fileStr = e.file
	}
	if e.line != 0 {
		lineStr = strconv.Itoa(e.line)
	}
	if e.err != nil {
		errStr = e.err.Error()
	}
	return fmt.Sprintf("Check failed at %s:%s (err:%s, args=%#v)", fileStr, lineStr, errStr, e.args)
}
func (e *checkerError) OrigError() error {
	return e.err
}
func (e *checkerError) Args() []interface{} {
	return e.args
}
func (e *checkerError) Location() (string, int) {
	return e.file, e.line
}
func (e *checkerError) StackTrace() []byte {
	return e.stackTrace
}
func (e *checkerError) Checker() Checker {
	return e.checker
}
