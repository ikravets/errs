package errs

type checkerLight struct {
	baseCallerDepth int
}

func NewCheckerLight(baseCallerDepth int) Checker {
	return &checkerLight{
		baseCallerDepth: baseCallerDepth,
	}
}

func (c *checkerLight) CheckE(err error, args ...interface{}) {
	if err != nil {
		panic(newCheckerError(c.baseCallerDepth+1, c, err, args))
	}
}
func (c *checkerLight) Check(cond bool, args ...interface{}) {
	if !cond {
		panic(newCheckerError(c.baseCallerDepth+1, c, nil, args))
	}
}
func (c *checkerLight) PassE(errptr *error) {
	if r := recover(); r != nil {
		c.PassFromRecover(errptr, r)
	}
}
func (c *checkerLight) PassFromRecover(errptr *error, recovered interface{}) {
	r := recovered
	ce, ok := r.(*checkerError)
	if !ok {
		// XXX no way to keep stack trace when re-panicing :(
		panic(r)
	}
	if errptr == nil {
		return
	}
	if ce.err != nil {
		*errptr = ce.err
	} else {
		*errptr = ce
	}
}
func (c *checkerLight) Catch(f func(CheckerError)) {
	if r := recover(); r != nil {
		c.CatchFromRecover(f, r)
	}
}
func (c *checkerLight) CatchFromRecover(f func(CheckerError), recovered interface{}) {
	ce, ok := recovered.(*checkerError)
	if !ok {
		// XXX no way to keep stack trace when re-panicing :(
		panic(recovered)
	}
	f(ce)
}
func (c *checkerLight) Is(checker Checker) bool {
	return checker.(*checkerLight) == c
}
