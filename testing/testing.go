package filterTesting

import (
	"testing"

	"github.com/flosch/pongo2"
)

// UnTemplate processes the template text, while handling test scenarios
func UnTemplate(t *testing.T, tpltxt string) (string, error) {
	tpl, err := pongo2.FromString(tpltxt)
	if err != nil {
		t.Fatalf("Error parsing template('%s'): %s", tpltxt, err.Error())
		return "", err
	}
	return tpl.Execute(pongo2.Context{})
}
