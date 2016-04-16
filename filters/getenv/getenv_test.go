package GetEnvFilter

import (
	"os"
	"testing"

	"github.com/odedlaz/untem/template"
)

func TestGetenv(t *testing.T) {
	os.Setenv("SET", "OK")
	defer os.Unsetenv("SET")
	txt, err := template.Execute("{{ \"SET\"|getenv }}", nil)
	if txt != "OK" || err != nil {
		t.Fail()
	}
}

func TestGetenvDefaultWhenMissing(t *testing.T) {
	txt, err := template.Execute("{{ \"NOT_SET\"|getenv:\"OK\" }}", nil)
	if txt != "OK" || err != nil {
		t.Fail()
	}
}

func TestGetenvMissingFails(t *testing.T) {
	_, err := template.Execute("{{ \"NOT_SET\"|getenv }}", nil)
	if err == nil {
		t.Fail()
	}
}
