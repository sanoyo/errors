package errors

import "runtime"

func New(message string) error {
	return &base{
		message: message,
		stack:   callers(),
	}
}

type base struct {
	message string
	*stack
}

func (b *base) Error() string {
	return b.message
}

type stack []uintptr

func callers() *stack {
	const depth = 16
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
