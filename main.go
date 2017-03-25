package main

import "log"

func init() {
	subCommands["help"] = &helpCmd
	subCommands["get-log"] = &getLogsCmd
	subCommands["get-logs"] = &getLogsCmd
	subCommands["bounce"] = &bounceCmd
}

func main() {
	log.SetFlags(0)
	opts := parseOpts()

	runCommand(opts)
}

func runCommand(opts *options) {
	if cmd, found := subCommands[opts.cmd]; found {
		cmd.run(opts.cmd, opts)
		return
	}
	log.Fatalf("Subcommand not recognized: %s", opts.cmd)
}
