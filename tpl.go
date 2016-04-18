package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/odedlaz/tpl/config"
	// load all the filters
	_ "github.com/odedlaz/tpl/filters"
	tplos "github.com/odedlaz/tpl/os"
	tplpath "github.com/odedlaz/tpl/os/path"
	tpl "github.com/odedlaz/tpl/template"
)

// VERSION current version of the app
const VERSION = "0.2-dev"

// default config path

func printVersion(ctx *kingpin.ParseContext) error {
	fmt.Printf("tpl %s\n", VERSION)
	fmt.Printf("filters: %s\n", strings.Join(tpl.Filters, ", "))
	os.Exit(0)
	return nil
}

func handleUnhandledError() {
	if err := recover(); err != nil {
		app.Fatalf(fmt.Sprint(err))
	}
}

func editInPlace(text, templateText string) {
	mode, err := tplpath.FileMode(*templateFilename)
	if err != nil {
		app.Fatalf("edit file (%s) in place: %v, try --help", *configFilename, err.Error())
	}

	filesToWrite := map[string]string{
		fmt.Sprintf("%s.bak", *templateFilename): templateText,
		*templateFilename:                        text}

	for filename, txt := range filesToWrite {
		ioutil.WriteFile(filename,
			[]byte(txt),
			mode)
	}
}

var (
	app              = kingpin.New("tpl", "A command-line un-templating application.")
	edit             = app.Flag("in-place", "edit files in place (if --filename was supplied)").Short('i').Bool()
	templateFilename = app.Arg("filename", "path to a file containing a valid djanto template (DTL)").ExistingFile()
	configFilename   = app.Flag("config", "path to config file").Default(config.DEFAULT_FILENAME).String()
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
	} else if !os.IsNotExist(err) || *configFilename != config.DEFAULT_FILENAME {
		app.Fatalf("loading config from: '%s': %v, try --help", *configFilename, err.Error())
	}

	templateText := tpl.Must(tpl.Read(*templateFilename))

	text := tpl.Must(tpl.Execute(templateText))

	if *edit {
		editInPlace(text, templateText)
		return
	}
	fmt.Print(text)
}
