package main

import (
	"flag"
	"fmt"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	var ops Ops
	flags.StringVar(&ops.Config, "config", "", "Pox config path")
	flags.StringVar(&ops.Config, "c", "", "Pox config path(Short)")
	flags.StringVar(&ops.Root, "root", "/var/www/pox", "Pox response document root")
	flags.StringVar(&ops.Root, "r", "/var/www/pox", "Pox response document root(Short)")
	flags.BoolVar(&ops.Verbose, "verbose", false, "Print verbose log")
	flags.BoolVar(&ops.Verbose, "v", false, "Print verbose log(Short)")
	flags.IntVar(&ops.Delay, "delay", 1, "Delay seconds for response")
	flags.IntVar(&ops.Delay, "d", 1, "Delay seconds for response(Short)")
	flags.StringVar(&ops.Type, "type", "rest", "Api type")
	flags.StringVar(&ops.Type, "t", "rest", "Api type(Short)")
	flags.BoolVar(&ops.Version, "version", false, "Print version information and quit.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if ops.Version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	return Pox(ops)
}
