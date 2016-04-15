package GetEnvFilter

import (
	"os"
	"testing"

	filterTesting "github.com/odedlaz/untemplate-me/testing"
)

func TestGetenv(t *testing.T) {
	os.Setenv("SET", "OK")
	defer os.Unsetenv("SET")
	txt, err := filterTesting.UnTemplate(t, "{{ \"SET\"|getenv }}")
	if txt != "OK" || err != nil {
		t.Fail()
	}
}

func TestGetenvDefaultWhenMissing(t *testing.T) {
	txt, err := filterTesting.UnTemplate(t, "{{ \"NOT_SET\"|getenv:\"OK\" }}")
	if txt != "OK" || err != nil {
		t.Fail()
	}
}

func TestGetenvMissingFails(t *testing.T) {
	_, err := filterTesting.UnTemplate(t, "{{ \"NOT_SET\"|getenv }}")
	if err == nil {
		t.Fail()
	}
}
