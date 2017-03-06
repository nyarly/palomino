package main

import (
	"log"

	"github.com/SeeSpotRun/coerce"
	docopt "github.com/docopt/docopt-go"
)

type subCommand struct {
	description string
	docs        string
	opts        interface{}
	action      func(*options, interface{})
}

func (sc *subCommand) run(name string, opts *options) {
	argv := []string{name}
	argv = append(argv, opts.args...)
	sc.parseArgv(argv)

	sc.action(opts, sc.opts)
}

func fullDocs(docs string) string {
	return docs + `
For common options, see 'palomino help'
`
}

func (sc *subCommand) parseArgv(argv []string) {
	parsed, err := docopt.Parse(fullDocs(sc.docs), argv, true, "", false)

	if err != nil {
		log.Fatal(err)
	}

	err = coerce.Struct(sc.opts, parsed, "-%s", "--%s", "<%s>")
	if err != nil {
		log.Fatal(err)
	}
}

var subCommands = map[string]*subCommand{}
