package CatFilter

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"

	tplos "github.com/odedlaz/tpl/core/os"
	"github.com/odedlaz/tpl/template"
)

func init() {
	template.RegisterFilter("cat", cat)
}

func readAllFiles(paths []string) string {
	var buffer bytes.Buffer
	for _, path := range paths {
		txt, err := tplos.ReadFile(path)
		if err != nil {
			// skip invalid paths, but don't ignore them!
			fmt.Fprintf(os.Stderr, "cat: %s: %v\n", path, err.Error())
			continue
		}
		buffer.WriteString(fmt.Sprintln(txt))
	}

	return buffer.String()
}

func cat(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	path := in.String()
	paths, err := filepath.Glob(path)
	if err != nil {
		return nil, &pongo2.Error{
			Sender:   "filter:cat",
			ErrorMsg: fmt.Sprintf("error glob from '%s': %v", path, err.Error()),
		}
	}
	txt := readAllFiles(paths)

	if txt == "" {
		return nil, &pongo2.Error{
			Sender:   "filter:cat",
			ErrorMsg: fmt.Sprintf("no readable files at '%s'", path),
		}
	}
	return pongo2.AsValue(strings.TrimSuffix(txt, "\n")), nil
}
