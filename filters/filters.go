package filters

import (
	"github.com/flosch/pongo2"
	// getenv filter
	_ "github.com/odedlaz/untemplate-me/filters/getenv"
	// getenv filter
	_ "github.com/odedlaz/untemplate-me/filters/httpget"
)

// UnTemplate transforms a template text into whatever it needed to be
func UnTemplate(tpltxt string) (string, error) {
	tpl, err := pongo2.FromString(tpltxt)
	if err != nil {
		return "", err
	}
	return tpl.Execute(pongo2.Context{})
}
