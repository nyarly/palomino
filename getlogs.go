package main

import "log"

var getLogsCmd = subCommand{
	description: "retrieve logs for a deployment",
	docs:        glDocs,
	opts:        &glOpts{},
	action:      getLogs,
}

type glOpts struct {
	request, deploy string
}

const glDocs = `Retrieve logs
Usage: palomino get-log <request> [<deploy>]

`

func getLogs(opts *options, glif interface{}) {
	glopts := glif.(*glOpts)

	log.Printf("%#v", glopts)
}
