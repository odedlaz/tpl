package template

import (
	"errors"
	"regexp"
	"strings"

	"gopkg.in/flosch/pongo2.v3"

	"github.com/odedlaz/tpl/core/config"
	tplos "github.com/odedlaz/tpl/core/os"
)

// Filters all tpl registered filters
var Filters = []string{}

// Functions all tpl registered functions
var Functions = pongo2.Context{}

// Settings global settings
var (
	Settings = config.Settings{}
	// remove any empty lines (even if there are spaces between them)
	stripEmptyLinesRegex = regexp.MustCompile("\n\\s*\n\\s*\n")
)

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

	txt, err := tpl.Execute(Functions)
	if err != nil {
		return "", err
	}
	txt = stripEmptyLinesRegex.ReplaceAllString(txt, "\n")
	return strings.TrimLeft(txt, "\n"), nil
}

// RegisterFilter adds a new filter to pongo2
func RegisterFilter(name string, fn pongo2.FilterFunction) {
	pongo2.RegisterFilter(name, fn)
	Filters = append(Filters, name)
}

// RegisterFunction adds a new function to pongo2
func RegisterFunction(name string, fn interface{}) {
	Functions[name] = fn
}

// RegisterSettings adds a global settings to pongo
func RegisterSettings(settings *config.Settings) {
	Settings = *settings
}
