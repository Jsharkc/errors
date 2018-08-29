package errors

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

type errorWithStack struct {
	err   error
	stack string
}

func (e errorWithStack) Error() string {
	return fmt.Sprintf("error: %v\nstack: \n%s\n", e.err, e.stack)
}

// NewWithLevel create new error with stack level
func NewWithLevel(err string, level int32) errorWithStack {
	stack := string(debug.Stack())
	stacks := strings.Split(stack, "\n")

	if level > 0 {
		stack = strings.Join(stacks[5:5+level*2], "\n")
	} else {
		stack = strings.Join(stacks[5:], "\n")
	}

	er := errorWithStack{
		err:   errors.New(err),
		stack: stack,
	}

	return er
}

// WrapWithLevel wrap error with stack level
func WrapWithLevel(err error, level int32) errorWithStack {
	stack := string(debug.Stack())
	stacks := strings.Split(stack, "\n")

	if level > 0 {
		stack = strings.Join(stacks[5:5+level*2], "\n")
	} else {
		stack = strings.Join(stacks[5:], "\n")
	}

	er := errorWithStack{
		err:   err,
		stack: stack,
	}

	return er
}
