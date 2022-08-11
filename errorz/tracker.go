package errorz

import (
	"bytes"
	"io"
	"sync"
)

type Errors []error

type Tracker struct {
	errz Errors
}

// _bufferPool is a pool of bytes.Buffers.
var _bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func Track[S, D any](e *Tracker, call func(S) (D, error), input S) D {
	ret, err := call(input)
	if err != nil {
		e.errz = append(e.errz, err)
	}

	return ret
}

func (e *Tracker) Append(err error) {
	if err != nil {
		e.errz = append(e.errz, err)
	}
}

func (e *Tracker) Errors() Errors {
	return e.errz
}

func (e Errors) Error() string {
	buf := _bufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	e.doErrors(buf, 0)

	result := buf.String()
	_bufferPool.Put(buf)
	return result
}

func (e Errors) doErrors(w io.Writer, indent int) {
	for _, err := range e {
		for i := 0; i <= indent; i++ {
			io.WriteString(w, "  ")
		}
		io.WriteString(w, "- ")
		if eg, ok := err.(Errors); ok {
			eg.doErrors(w, indent+1)
		} else {
			io.WriteString(w, err.Error())
			io.WriteString(w, "\n")
		}
	}
}
