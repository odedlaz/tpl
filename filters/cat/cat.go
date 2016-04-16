package CatFilter

import (
	"fmt"
	"strings"

	"github.com/flosch/pongo2"

	tplos "github.com/odedlaz/tpl/os"
	"github.com/odedlaz/tpl/template"
)

func init() {
	template.RegisterFilter("cat", cat)
}

func cat(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	filename := in.String()
	txt, err := tplos.ReadFile(filename)

	if err != nil && param.IsNil() {
		return nil, &pongo2.Error{
			Sender:   "filter:cat",
			ErrorMsg: fmt.Sprintf("problem accessing '%s': %v", filename, err.Error()),
		}
	}

	if err == nil {
		return pongo2.AsValue(strings.TrimSuffix(txt, "\n")), nil
	}

	return pongo2.AsValue(param.String()), nil
}
