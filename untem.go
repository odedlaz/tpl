package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/odedlaz/untem/config"
	// load all the filters
	_ "github.com/odedlaz/untem/filters"
	"github.com/odedlaz/untem/flags"
	untemos "github.com/odedlaz/untem/os"
	"github.com/odedlaz/untem/template"
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
		return untemos.ReadFile(*filename)
	}

	// fallback - read from pipe
	return untemos.ReadFromStdIn(), nil
}

func Must(txt string, err error) string {
	if err != nil {
		panic(err)
	}
	return txt
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()

	flags.Parse()

	if *filename == "" && !untemos.StdInAvailable() {
		flags.Fatal("you have to pipe text or pass a filename argument")
		os.Exit(1)
	}

	settings, err := config.Load(*configFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't load configuration from '%s', defaulting to none.\n", *configFilename)
	}

	tplText := Must(getTemplateText())

	txt := Must(template.Execute(tplText, settings))

	fmt.Print(txt)
}
