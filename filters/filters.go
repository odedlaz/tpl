package filters

import (
	"github.com/flosch/pongo2"
	"github.com/odedlaz/untemplate-me/config"

	// getenv filter
	_ "github.com/odedlaz/untemplate-me/filters/getenv"
	// getenv filter
	_ "github.com/odedlaz/untemplate-me/filters/httpget"
	// kv filter
	_ "github.com/odedlaz/untemplate-me/filters/kv"
)

// UnTemplate transforms a template text into whatever it needed to be
func UnTemplate(tpltxt string, settings *config.Settings) (string, error) {
	tpl, err := pongo2.FromString(tpltxt)
	if err != nil {
		return "", err
	}

	pongo2.Globals["settings"] = settings

	return tpl.Execute(pongo2.Context{})
}
