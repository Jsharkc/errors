package errors

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

var DefaultStackLevel = 2

type errorWithStack struct {
	err   error
	stack string
}

func (e errorWithStack) Error() string {
	return fmt.Sprintf("%v\nstack: \n%s\n", e.err, e.stack)
}

// NewWithLevel create new error with stack level
func New(err string) errorWithStack {
	return NewWithLevel(err, 0)
}

// NewWithLevel create new error with stack level
func NewWithLevel(err string, level uint32) errorWithStack {
	stack := string(debug.Stack())
	stacks := strings.Split(stack, "\n")

	if level > 0 {
		stack = strings.Join(stacks[5:5+level*2], "\n")
	} else {
		stack = strings.Join(stacks[5:5+DefaultStackLevel*2], "\n")
	}

	er := errorWithStack{
		err:   errors.New(err),
		stack: stack,
	}

	return er
}

// Wrap wrap error with stack level
func Wrap(err error) errorWithStack {
	return WrapWithLevel(err, 0)
}

// WrapWithLevel wrap error with stack level
func WrapWithLevel(err error, level uint32) errorWithStack {
	stack := string(debug.Stack())
	stacks := strings.Split(stack, "\n")

	if level > 0 {
		stack = strings.Join(stacks[5:5+level*2], "\n")
	} else {
		stack = strings.Join(stacks[5:5+DefaultStackLevel*2], "\n")
	}

	er := errorWithStack{
		err:   err,
		stack: stack,
	}

	return er
}
