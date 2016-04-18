package CatFilter

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

func TestCat(t *testing.T) {
	tpl := fmt.Sprintf("{{ \"%s\" | cat }}", tmpfile.Name())
	txt, err := template.Execute(tpl)
	if txt != "OK" || err != nil {
		t.Fail()
	}
}

func TestCatMissingFails(t *testing.T) {
	tpl := "{{ \"/no/path\" | cat }}"
	txt, err := template.Execute(tpl)
	if err == nil {
		t.Errorf("got a test, altough we shouldn't have: %s", txt)
	}
}
