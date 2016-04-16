package HttpGetFilter

import (
	"testing"

	"github.com/odedlaz/untem/template"
)

func TestHttpGet(t *testing.T) {
	txt, err := template.Execute("{{ \"https://api.ipify.org\" | httpget }}", nil)
	if txt == "" || err != nil {
		t.Fail()
	}
}

func TestHttpGetInvalidUrl(t *testing.T) {
	_, err := template.Execute("{{ \"invalid\" | httpget }}", nil)
	if err == nil {
		t.Errorf("An error should have been thrown when accessing an invalid url: %v", err.Error())
	}
}
