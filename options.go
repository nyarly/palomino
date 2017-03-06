package main

import (
	"fmt"
	"log"

	"github.com/SeeSpotRun/coerce"
	docopt "github.com/docopt/docopt-go"
)

type options struct {
	cmd  string
	args []string
	url  string
}

const docstring = `Issue commands against a Singularity
Usage: palomino [options] <cmd> [<args>...]

Options:
  --url=<string>  The URL of the singularity to contact
`

func docStr() string {
	str := docstring + "\nMost common subcommands:"
	for _, k := range []string{"help", "get-log"} {
		str = str + fmt.Sprintf("\n  %s: %s",
			k, subCommands[k].description)
	}
	str = str + "\n"
	return str
}

func parseArgv(argv []string) *options {
	parsed, err := docopt.Parse(docStr(), argv, true, "palomino v0.0.1", false)

	if err != nil {
		log.Fatal(err)
	}

	opts := options{}
	err = coerce.Struct(&opts, parsed, "-%s", "--%s", "<%s>")
	if err != nil {
		log.Fatal(err)
	}

	return &opts
}

func parseOpts() *options {
	opts := parseArgv(nil)
	processOpts(opts)

	return opts
}

func processOpts(opts *options) {
	if opts.cmd == "" {
		opts.cmd = "help"
	}
}
