package cerrors

import (
	"fmt"
	"io"
)

type cause struct {
	cause error
	msg   string
	code  int
}

func (c cause) Error() string {
	if len(c.msg) == 0 {
		return fmt.Sprintf("Error code %v: %v", c.code, c.Cause())
	} else {
		return fmt.Sprintf("Error code %v. %s: %v", c.code, c.msg, c.Cause())
	}
}
func (c cause) Cause() error { return c.cause }

// wrapper is an error implementation returned by Wrap and Wrapf
// that implements its own fmt.Formatter.
type wrapper struct {
	cause
	*stack
}

func (w wrapper) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			fmt.Fprintf(s, "Error code %v.", w.code)
			if len(w.msg) != 0 {
				io.WriteString(s, "\n"+w.msg+".")
			}
			fmt.Fprintf(s, "%+v", w.StackTrace())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

// Wrap returns an error annotating err with message.
// If err is nil, Wrap returns nil.
func Wrap(err error, code int, message string) error {
	if err == nil {
		return nil
	}
	return wrapper{
		cause: cause{
			cause: err,
			msg:   message,
			code:  code,
		},
		stack: callers(),
	}
}

// Wrapf returns an error annotating err with the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrapper{
		cause: cause{
			cause: err,
			msg:   fmt.Sprintf(format, args...),
			code:  code,
		},
		stack: callers(),
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
