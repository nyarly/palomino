package main

var helpCmd = subCommand{
	description: "prints help",
	docs:        helpDocs,
	opts:        &helpOpts{},
}

func init() {
	helpCmd.action = help
}

type helpOpts struct {
	cmd string
}

const helpDocs = `get help
Usage: palomino help [<cmd>]
`

func help(mainOpts *options, helpif interface{}) {
	opts := helpif.(*helpOpts)

	switch opts.cmd {
	default:
		mainOpts.cmd = opts.cmd
		mainOpts.args = []string{"--help"}
		runCommand(mainOpts)
	case "":
		parseArgv([]string{"--help"})
	case "help":
		helpCmd.parseArgv([]string{"--help"})
	}
}
