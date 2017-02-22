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

func (sc *subCommand) run(opts *options) {
	sc.parseOpts()

	sc.action(opts, sc.opts)
}

func (sc *subCommand) parseOpts() {
	sc.parseArgv(nil)
}

func (sc *subCommand) parseArgv(argv []string) {
	parsed, err := docopt.Parse(sc.docs, argv, true, "", false)

	if err != nil {
		log.Fatal(err)
	}

	err = coerce.Struct(sc.opts, parsed, "-%s", "--%s", "<%s>")
	if err != nil {
		log.Fatal(err)
	}
}

var subCommands = map[string]*subCommand{}
