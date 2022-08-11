package errorz

import (
	"go.uber.org/multierr"
)

type Tracker struct {
	errz error
}

func Must[S, D any](e *Tracker, call func(S) (D, error), input S) D {
	ret, err := call(input)
	if err != nil {
		e.errz = multierr.Append(e.errz, err)
	}

	return ret
}

type ErrorGroup interface {
	Errors() []error
}

type Errors []error

func (e *Tracker) Errors() Errors {
	if e.errz == nil {
		return nil
	}
	if eg, ok := e.errz.(ErrorGroup); ok {
		return eg.Errors()
	}
	return nil
}
