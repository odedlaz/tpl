package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/odedlaz/untemplate-me/filters"
	"github.com/odedlaz/untemplate-me/flags"
	untemos "github.com/odedlaz/untemplate-me/os"
)

var (
	filename = flags.App.Arg("filename", "path to a file containing a valid djanto template (DTL)").ExistingFile()
)

func getTemplateText() (string, error) {
	if *filename == "" && !untemos.StdInAvailable() {
		return "", errors.New("you have to pipe text or pass a filename argument")
	}

	// if a filename was passed - use it instead of the pipe
	if *filename != "" {
		return untemos.ReadFromFile(*filename)
	}

	// fallback - read from pipe
	return untemos.ReadFromStdIn(), nil
}

func main() {
	flags.Parse()
	tplText, err := getTemplateText()
	if err != nil {
		flags.Fatal(err.Error())
		os.Exit(1)
	}
	txt, err := filters.UnTemplate(tplText)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Print(txt)
}
