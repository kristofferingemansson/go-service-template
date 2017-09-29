package pkg

import (
	"fmt"
	"os"
)

// Logger ..
type Logger interface {
	Log(keyvals ...interface{})
}

type logger struct{}

func (l logger) Log(keyvals ...interface{}) {
	fmt.Fprintln(os.Stdout, keyvals...)
}

// StdLogger ..
var StdLogger Logger = logger{}
