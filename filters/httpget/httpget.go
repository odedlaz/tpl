package HttpGetFilter

import (
	"github.com/flosch/pongo2"
	"github.com/franela/goreq"
)

func init() {
	pongo2.RegisterFilter("httpget", httpGet)
}

func httpGet(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	url := in.String()
	response, err := goreq.Request{Uri: url}.Do()

	if err != nil {
		return nil, &pongo2.Error{
			Sender:   "filter:httpget",
			ErrorMsg: err.Error(),
		}
	}

	content, err := response.Body.ToString()
	if err != nil {
		return nil, &pongo2.Error{
			Sender:   "filter:httpget",
			ErrorMsg: err.Error(),
		}
	}

	return pongo2.AsValue(content), nil
}
