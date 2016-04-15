package main

import (
	"errors"
	"fmt"
	"os"

	conf "github.com/odedlaz/untemplate-me/config"
	"github.com/odedlaz/untemplate-me/filters"
	"github.com/odedlaz/untemplate-me/flags"
	untemos "github.com/odedlaz/untemplate-me/os"
)

var (
	filename       = flags.App.Arg("filename", "path to a file containing a valid djanto template (DTL)").ExistingFile()
	configFilename = flags.App.Flag("config", "path to config file").Default("untem.yaml").String()
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

	settings, err := conf.Load(*configFilename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	txt, err := filters.UnTemplate(tplText, settings)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
	fmt.Print(txt)
}
