package PathExistsFilter

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/odedlaz/tpl/template"
)

var tmpfile *os.File

// TODO: mock file system calls.
// maybe use: https://github.com/blang/vfs
func setup() {
	tmpfile, _ = ioutil.TempFile("/tmp", "tpl")
	tmpfile.Write([]byte("OK"))
	tmpfile.Close()
}

func teardown() {
	os.Remove(tmpfile.Name())
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func TestPathExists(t *testing.T) {
	tpl := fmt.Sprintf("{%% if \"%s\" | pathexists %%}OK{%% endif %%}", tmpfile.Name())
	txt, err := template.Execute(tpl)
	if txt != "OK" || err != nil {
		t.Fail()
	}
}

func TestPathDoesNotExist(t *testing.T) {
	tpl := "{% if not \"/does/not/exist\" | pathexists %}OK{% endif %}"
	txt, err := template.Execute(tpl)
	if txt != "OK" || err != nil {
		t.Fail()
	}
}
