package main

import "log"

func init() {
	subCommands["help"] = &helpCmd
	subCommands["get-log"] = &getLogsCmd
}

func main() {
	log.SetFlags(log.Lshortfile)
	opts := parseOpts()

	runCommand(opts)
}

func runCommand(opts *options) {
	if cmd, found := subCommands[opts.cmd]; found {
		cmd.run(opts)
		return
	}
	log.Fatal("Subcommand not recognized: %s", opts.cmd)
}
