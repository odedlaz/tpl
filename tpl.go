package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/odedlaz/tpl/config"
	// load all the filters
	_ "github.com/odedlaz/tpl/filters"
	tplos "github.com/odedlaz/tpl/os"
	tpl "github.com/odedlaz/tpl/template"
)

// VERSION current version of the app
const VERSION = "0.1"

func printVersion(ctx *kingpin.ParseContext) error {
	fmt.Printf("tpl %s\n", VERSION)
	fmt.Printf("filters: %s\n", strings.Join(tpl.Filters, ", "))

	os.Exit(0)
	return nil
}

var (
	app              = kingpin.New("tpl", "A command-line un-templating application.")
	templateFilename = app.Arg("filename", "path to a file containing a valid djanto template (DTL)").ExistingFile()
	configFilename   = app.Flag("config", "path to config file").Default("tpl.yaml").String()
	showVersion      = app.Flag("version", "show version and quit").Action(printVersion).Bool()
)

func main() {
	// fail gracefully
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()

	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *templateFilename == "" && !tplos.StdInAvailable() {
		app.Fatalf("you have to pipe text or pass a template filename as argument, try --help")
		os.Exit(1)
	}

	// if settings can't be loaded, the default is used
	if settings, err := config.Load(*configFilename); err != nil {
		fmt.Fprintf(os.Stderr,
			"couldn't load configuration from '%s', ignoring configuration.\n",
			*configFilename)
	} else {
		tpl.RegisterSettings(settings)
	}

	templateText := tpl.Must(tpl.Read(*templateFilename))

	text := tpl.Must(tpl.Execute(templateText))

	fmt.Print(text)
}
