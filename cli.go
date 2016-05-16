package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %s [options]\n\nOptions:\n", Name)
		flags.VisitAll(func(f *flag.Flag) {
			if len(f.Name) == 1 {
				s := fmt.Sprintf("  -%s", f.Name)
				fmt.Fprint(os.Stderr, s, ",")
			} else {
				s := fmt.Sprintf(" --%s", f.Name)
				_, usage := flag.UnquoteUsage(f)
				num := 12 - len(f.Name)
				s += strings.Repeat(" ", num) + usage
				if !(f.DefValue == "" || f.DefValue == "false" || f.DefValue == "0") {
					s += fmt.Sprintf(" (default %v)", f.DefValue)
				}
				fmt.Fprint(os.Stderr, s, "\n")
			}
		})
	}

	conf := os.Getenv(strings.ToUpper(Name) + "_CONF")
	c := DefaultConfig()

	var ops Ops
	flags.StringVar(&ops.Config, "config", conf, "config path")
	flags.StringVar(&ops.Config, "c", conf, "")

	flags.StringVar(&ops.Root, "root", c.Root, "document root")
	flags.StringVar(&ops.Root, "r", c.Root, "")

	flags.StringVar(&ops.Addr, "addr", c.Addr, "address with port")
	flags.StringVar(&ops.Addr, "a", c.Addr, "")

	flags.StringVar(&ops.LogLevel, "log-level", c.LogLevel, "log level")
	flags.StringVar(&ops.LogLevel, "l", c.LogLevel, "")

	flags.IntVar(&ops.Delay, "delay", c.Delay, "delay seconds for response")
	flags.IntVar(&ops.Delay, "d", c.Delay, "")

	flags.StringVar(&ops.Protocol, "protocol", c.Protocol, "api protocol")
	flags.StringVar(&ops.Protocol, "p", c.Protocol, "")

	flags.BoolVar(&ops.Version, "version", false, "print the version and exit")
	flags.BoolVar(&ops.Version, "v", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if ops.Version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	return Pox(ops)
}
