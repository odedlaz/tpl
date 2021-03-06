package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	// load all the filters
	_ "github.com/odedlaz/tpl/filters"
	// load all the functions
	_ "github.com/odedlaz/tpl/functions"

	"github.com/odedlaz/tpl/core"
	"github.com/odedlaz/tpl/core/config"
	tplenum "github.com/odedlaz/tpl/core/enumerable"
	tplos "github.com/odedlaz/tpl/core/os"
	tplpath "github.com/odedlaz/tpl/core/os/path"
	tpl "github.com/odedlaz/tpl/template"
)

// default config path

func printVersion(ctx *kingpin.ParseContext) error {
	fmt.Printf("tpl %s\nFilters: %s\nFunctions: %s\n",
		core.Version,
		strings.Join(tpl.Filters, " "),
		strings.Join(tplenum.MapKeys(tpl.Functions), " "))
	os.Exit(0)
	return nil
}

func handleUnhandledError() {
	if err := recover(); err != nil {
		app.Fatalf(fmt.Sprint(err))
	}
}

func editInPlace(text string) (err error) {
	var mode os.FileMode

	if mode, err = tplpath.FileMode(*templateFilename); err != nil {
		return err
	}

	if err = os.Rename(*templateFilename, fmt.Sprintf("%s.bak", *templateFilename)); err != nil {
		return err
	}

	if err = ioutil.WriteFile(*templateFilename, []byte(text), mode); err != nil {
		return err
	}

	return nil
}

var (
	app              = kingpin.New("tpl", "A command-line un-templating application.")
	edit             = app.Flag("in-place", "edit files in place (if --filename was supplied)").Short('i').Bool()
	templateFilename = app.Arg("filename", "path to a file containing a valid djanto template (DTL)").ExistingFile()
	configFilename   = app.Flag("config", "path to config file").Default(config.DefaultFilename).String()
	showVersion      = app.Flag("version", "show version and quit").Action(printVersion).Bool()
)

func main() {
	defer handleUnhandledError()

	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *templateFilename == "" && !tplos.StdInAvailable() {
		app.Fatalf("you have to pipe text or pass a template filename as argument, try --help")
	}

	if *edit && *templateFilename == "" {
		app.Fatalf("you can't edit a file if you're piping the template, try --help")
	}

	// if settings can't be loaded, the default is used
	// if the file doesn't exit -> ignore
	if settings, err := config.Load(*configFilename); err == nil {
		tpl.RegisterSettings(settings)
	} else if !os.IsNotExist(err) || *configFilename != config.DefaultFilename {
		app.Fatalf("loading config from: '%s': %v, try --help", *configFilename, err.Error())
	}

	templateText := tpl.Must(tpl.Read(*templateFilename))

	text := tpl.Must(tpl.Execute(templateText))

	if !*edit {
		fmt.Print(text)
		return
	}

	if err := editInPlace(text); err != nil {
		app.Fatalf("edit file (%s) in place: %v, try --help", *templateFilename, err)
	}
}
