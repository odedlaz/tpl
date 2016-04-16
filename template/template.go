package template

import (
	"github.com/flosch/pongo2"
	"github.com/odedlaz/untem/config"
)

// Execute transforms a template text into whatever it needed to be
func Execute(tpltxt string, settings *config.Settings) (string, error) {
	tpl, err := pongo2.FromString(tpltxt)
	if err != nil {
		return "", err
	}

	pongo2.Globals["settings"] = settings

	return tpl.Execute(pongo2.Context{})
}
