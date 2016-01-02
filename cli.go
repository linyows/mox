package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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

	conf := os.Getenv("POX_CONF")
	c := DefaultConfig()

	var ops Ops
	flags.StringVar(&ops.Config, "config", conf, "Pox config path")
	flags.StringVar(&ops.Config, "c", conf, "Pox config path(Short)")

	flags.StringVar(&ops.Root, "root", c.Root, "Pox response document root")
	flags.StringVar(&ops.Root, "r", c.Root, "Pox response document root(Short)")

	flags.StringVar(&ops.Addr, "addr", c.Addr, "Server address with port")
	flags.StringVar(&ops.Addr, "a", c.Addr, "Server address with port(Short)")

	flags.StringVar(&ops.Loglevel, "loglevel", c.Loglevel, "Log level")
	flags.StringVar(&ops.Loglevel, "l", c.Loglevel, "Log level(Short)")

	flags.IntVar(&ops.Delay, "delay", c.Delay, "Delay seconds for response")
	flags.IntVar(&ops.Delay, "d", c.Delay, "Delay seconds for response(Short)")

	flags.StringVar(&ops.Type, "type", c.Type, "Api type")
	flags.StringVar(&ops.Type, "t", c.Type, "Api type(Short)")

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
