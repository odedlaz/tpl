package template

import (
	"errors"

	"github.com/flosch/pongo2"
	"github.com/odedlaz/tpl/config"

	tplos "github.com/odedlaz/tpl/os"
)

// Filters all tpl registered filters
var Filters = []string{}

// Settings global settings
var Settings = config.Settings{}

// Must panics if err != nil. otherwise, returns the text
func Must(txt string, err error) string {
	if err != nil {
		panic(err)
	}
	return txt
}

// Read reads the template from a file,
// or if it doesn't exit, tries to read it from stdin
func Read(filename string) (string, error) {
	if filename == "" && !tplos.StdInAvailable() {
		return "", errors.New("you have to pipe text or pass a filename argument")
	}

	// if a filename was passed - use it instead of the pipe
	if filename != "" {
		return tplos.ReadFile(filename)
	}

	// fallback - read from pipe
	return tplos.ReadFromStdIn(), nil
}

// Execute transforms a template text into whatever it needed to be
func Execute(tpltxt string) (string, error) {
	tpl, err := pongo2.FromString(tpltxt)
	if err != nil {
		return "", err
	}

	return tpl.Execute(pongo2.Context{})
}

// RegisterFilter adds a new filter to pongo2
func RegisterFilter(name string, fn pongo2.FilterFunction) {
	pongo2.RegisterFilter(name, fn)
	Filters = append(Filters, name)
}

// RegisterSettings adds a global settings to pongo
func RegisterSettings(settings *config.Settings) {
	Settings = *settings
}
