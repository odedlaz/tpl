package StringSplitFilter

import (
	"strings"

	"github.com/flosch/pongo2"

	"github.com/odedlaz/tpl/template"
)

func init() {
	template.RegisterFilter("stringsplit", split)
}

func split(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	if param.IsNil() {
		return nil, &pongo2.Error{
			Sender:   "filter:stringsplit",
			ErrorMsg: "you have to pass the delimiter as a parameter, for example: {{ \"1,2,3\" | stringsplit:\",\" }}",
		}
	}

	list := strings.Split(in.String(), param.String())
	return pongo2.AsValue(list), nil
}
