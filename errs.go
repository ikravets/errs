package errs

type CheckerError interface {
	error
	OrigError() error
	Args() []interface{}
	Location() (file string, line int)
	StackTrace() []byte
	Checker() Checker
}

type Checker interface {
	CheckE(err error, v ...interface{})
	Check(cond bool, v ...interface{})
	PassE(errptr *error)
	PassFromRecover(errptr *error, recovered interface{})
	Is(Checker) bool
	Catch(func(CheckerError))
	CatchFromRecover(f func(CheckerError), recovered interface{})
}

var DefaultChecker = NewCheckerLight(1)

func CheckE(err error, args ...interface{}) {
	DefaultChecker.CheckE(err, args...)
}
func Check(cond bool, args ...interface{}) {
	DefaultChecker.Check(cond, args...)
}
func PassE(errptr *error) {
	if r := recover(); r != nil {
		DefaultChecker.PassFromRecover(errptr, r)
	}
}
func Catch(f func(CheckerError)) {
	if r := recover(); r != nil {
		DefaultChecker.CatchFromRecover(f, r)
	}
}
