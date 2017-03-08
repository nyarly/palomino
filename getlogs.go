package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	singularity "github.com/opentable/go-singularity"
	"github.com/opentable/go-singularity/dtos"
	"github.com/opentable/swaggering"
)

var getLogsCmd = subCommand{
	description: "retrieve logs for a deployment",
	docs:        glDocs,
	opts:        &glOpts{},
	action:      getLogs,
}

type glOpts struct {
	request, deploy string
	options
}

const glDocs = `Retrieve logs
Usage: palomino [options] get-log <request> [<deploy>]
`

func getLogs(topOpts *options, glif interface{}) {
	opts := glif.(*glOpts)
	opts.options = *topOpts

	url, err := url.Parse(opts.url)
	if err != nil {
		log.Fatal(err)
	}

	var client *singularity.Client
	if opts.debug {
		client = singularity.NewClient(opts.url, swaggering.StdlibDebugLogger{})
	} else {
		client = singularity.NewClient(opts.url)
	}

	res, err := client.GetS3LogsForDeploy(opts.request, opts.deploy, 0, -1)
	if len(res) == 0 {
		log.Print("No S3 logs for request, attempting direct file download...")
		res, err := client.GetTaskHistoryForActiveRequest(opts.request)
		if err != nil {
			log.Print(err)
		}

		wait := &sync.WaitGroup{}
		for _, r := range res {
			task, err := client.GetActiveTask(r.TaskId.Id)
			if err != nil {
				log.Fatal(err)
			}
			if opts.deploy != "" {
				if opts.deploy != task.TaskRequest.Deploy.Id {
					continue
				}
			}

			box, err := client.Browse(r.TaskId.Id, "")
			if err != nil {
				log.Fatal(err)
			}
			//log.Printf("%d %#v", n, box)

			for _, f := range box.Files {
				wait.Add(1)
				go fetchFile(opts, task, *f, box, url, wait)
			}
		}

		wait.Wait()

	}
}

func fetchFile(opts *glOpts, task *dtos.SingularityTask, f dtos.SingularitySandboxFile, box *dtos.SingularitySandbox, url *url.URL, wait *sync.WaitGroup) {
	defer wait.Done()
	outPath := filepath.Join(os.TempDir(), opts.request, task.TaskRequest.Deploy.Id, f.Name)
	log.Printf("Writing to %q...", outPath)
	fetchURL := fmt.Sprintf("%s://%s:%d/files/download.json?path=%s/%s/%s", url.Scheme, box.SlaveHostname, opts.mesosPort, box.FullPathToRoot, box.CurrentDirectory, f.Name)
	http := &http.Client{}
	res, err := http.Get(fetchURL)
	if err != nil {
		log.Print(url, err)
		return
	}

	os.MkdirAll(filepath.Dir(outPath), os.ModePerm)

	file, err := os.Create(outPath)
	if err != nil {
		log.Print(outPath, err)
		return
	}

	wrote, err := io.Copy(file, res.Body)
	if err != nil {
		log.Print(outPath, err)
		return
	}
	log.Printf("... wrote %d bytes for %s:%s", wrote, task.TaskRequest.Deploy.Id, f.Name)
}
