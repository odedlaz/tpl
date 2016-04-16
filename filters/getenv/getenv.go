package GetEnvFilter

import (
	"fmt"
	"os"

	"github.com/flosch/pongo2"
	"github.com/odedlaz/tpl/template"
)

func init() {
	template.RegisterFilter("getenv", getenv)
}

func getenv(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	envName := in.String()
	envValue := os.Getenv(envName)

	if len(envValue) == 0 && param.IsNil() {
		return nil, &pongo2.Error{
			Sender:   "filter:getenv",
			ErrorMsg: fmt.Sprintf("'%s' isn't exported and no default variable given", envName),
		}
	}

	if len(envValue) > 0 {
		return pongo2.AsValue(envValue), nil
	}

	return pongo2.AsValue(param.String()), nil
}
