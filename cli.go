package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	flag "github.com/linyows/mflag"
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

// Ops structure
type Ops struct {
	Config   string
	Root     string
	Protocol string
	Addr     string
	Delay    int
	LogLevel string
	Version  bool
}

var usageText = `
Usage: mox [options]

Options:`

var exampleText = `
Example:
  $ mox --root /var/www/mox --protocol json-rpc --delay 1 --log-level debug
  $ mox --config /etc/mox.conf
`

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.outStream)

	flags.Usage = func() {
		fmt.Fprintf(cli.outStream, usageText)
		flags.PrintDefaults()
		fmt.Fprint(cli.outStream, exampleText)
	}

	conf := os.Getenv(strings.ToUpper(Name) + "_CONF")
	c := DefaultConfig()

	var ops Ops
	flags.StringVar(&ops.Config, []string{"c", "-config"}, conf, "config path")
	flags.StringVar(&ops.Root, []string{"r", "-root"}, c.Root, "document root path")
	flags.StringVar(&ops.Addr, []string{"a", "-addr"}, c.Addr, "network address with port")
	flags.StringVar(&ops.LogLevel, []string{"l", "-log-level"}, c.LogLevel, "log level")
	flags.IntVar(&ops.Delay, []string{"d", "-delay"}, c.Delay, "delay seconds for response")
	flags.StringVar(&ops.Protocol, []string{"p", "-protocol"}, c.Protocol, "api protocol -- REST or JSON-RPC")
	flags.BoolVar(&ops.Version, []string{"v", "-version"}, false, "print the version and exit")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if ops.Version {
		fmt.Fprintf(cli.outStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	return Pox(ops)
}
