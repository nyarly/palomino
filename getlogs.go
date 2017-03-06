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
Usage: palomino [options] get-log <request> [<deploy>]
`

func getLogs(opts *options, glif interface{}) {
	glopts := glif.(*glOpts)

	log.Printf("%#v", opts)
	log.Printf("%#v", glopts)
}
