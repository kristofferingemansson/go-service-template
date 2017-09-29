package pkg

import (
	"fmt"
	"io"
	"os"
)

// Logger ..
type Logger interface {
	Log(keyvals ...interface{})
}

type logger struct{ io.Writer }

func (l logger) Log(keyvals ...interface{}) {
	if l.Writer != nil {
		fmt.Fprintln(l.Writer, keyvals...)
	}
}

// StdLogger ..
var StdLogger Logger = logger{os.Stdout}

// NilLogger ..
var NilLogger Logger = logger{nil}
