package errors

import "runtime"

type base struct {
	message string
	*stack
}

func New(message string) error {
	return &base{
		message: message,
		stack:   callers(3),
	}
}

func Errorf(err error, message string) error {
	return &base{
		message: message + ": " + err.Error(),
		stack:   callers(3),
	}
}

type withStack struct {
	error
	*stack
}

func WithStack(err error) *withStack {
	if err == nil {
		return nil
	}
	return &withStack{
		err,
		callers(3),
	}
}

func (w *withStack) Cause(e error) error {
	return w.error
}

func (b *base) Error() string {
	return b.message
}

type stack []uintptr

func callers(skip int) *stack {
	const depth = 16
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
