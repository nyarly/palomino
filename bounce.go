package main

import (
	"log"

	singularity "github.com/opentable/go-singularity"
	"github.com/opentable/go-singularity/dtos"
	"github.com/opentable/swaggering"
)

var bounceCmd = subCommand{
	description: "bounce a request",
	docs:        bounceDocs,
	opts:        &bounceOpts{},
	action:      bounce,
}

type bounceOpts struct {
	options
	request string
}

const bounceDocs = `Bounce a request
Usage: palomino [options] bounce <request>
`

func bounce(topOpts *options, bif interface{}) {
	opts := bif.(*bounceOpts)
	opts.options = *topOpts

	var client *singularity.Client
	if opts.debug {
		client = singularity.NewClient(opts.url, swaggering.StdlibDebugLogger{})
	} else {
		client = singularity.NewClient(opts.url)
	}

	bounceReq, err := swaggering.LoadMap(&dtos.SingularityBounceRequest{}, map[string]interface{}{
		"Message":     "Palomino bounced",
		"Incremental": true,
	})

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Bounce(opts.request, bounceReq.(*dtos.SingularityBounceRequest))

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Done")
}
