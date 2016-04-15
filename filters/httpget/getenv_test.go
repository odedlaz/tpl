package HttpGetFilter

import (
	"testing"

	filterTesting "github.com/odedlaz/untemplate-me/testing"
)

func TestHttpGet(t *testing.T) {
	txt, err := filterTesting.UnTemplate(t, "{{ \"https://api.ipify.org\" | httpget }}")
	if txt == "" || err != nil {
		t.Fail()
	}
}

func TestHttpGetInvalidUrl(t *testing.T) {
	_, err := filterTesting.UnTemplate(t, "{{ \"invalid\" | httpget }}")
	if err == nil {
		t.Errorf("An error should have been thrown when accessing an invalid url: %v", err.Error())
	}
}
