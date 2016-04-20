package PathExistsFilter

import (
	"fmt"
	"os"

	"gopkg.in/flosch/pongo2.v3"

	"github.com/odedlaz/tpl/template"
)

func init() {
	template.RegisterFilter("pathexists", exists)
}

func exists(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	_, err := os.Stat(in.String())

	if err != nil && !os.IsNotExist(err) {
		return nil, &pongo2.Error{
			Sender:   "filter:osexists",
			ErrorMsg: fmt.Sprintf("error while checking of path existence at '%s': %v", in.String(), err.Error()),
		}
	}

	// the only way we got here is either
	// the err is nil which means the path exists
	// or os.IsNotExist(err) is true, which means the path doesn't exist
	return pongo2.AsValue(err == nil), nil
}
