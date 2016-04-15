package flags

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var App = kingpin.New("untemplate-me", "A command-line un-templating application.")

func Parse() {
	kingpin.MustParse(App.Parse(os.Args[1:]))
}

func Fatal(format string, args ...interface{}) {
	format = fmt.Sprintf("%s, try --help", format)
	kingpin.Fatalf(format, args...)
}
