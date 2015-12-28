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
	var (
		c string
		r string
		v bool
		d int
		t string

		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&c, "config", "", "Pox config path")
	flags.StringVar(&c, "c", "", "Pox config path(Short)")
	flags.StringVar(&c, "root", "/var/www/pox", "Pox response document root")
	flags.StringVar(&c, "r", "/var/www/pox", "Pox response document root(Short)")
	flags.BoolVar(&v, "v", false, "Print verbose log")
	flags.BoolVar(&v, "v", false, "Print verbose log(Short)")
	flags.IntVar(&d, "delay", 1, "Delay seconds for response")
	flags.IntVar(&d, "d", 1, "Delay seconds for response(Short)")
	flags.StringVar(&t, "type", "rest", "Api type")
	flags.StringVar(&t, "t", "rest", "Api type(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	_ = c
	_ = r
	_ = v
	_ = d
	_ = t

	return ExitCodeOK
}
