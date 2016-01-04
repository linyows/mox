package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/logutils"
)

// Pox starts server.
func Pox(ops Ops) int {
	c := DefaultConfig()
	c.Set(ops)

	if ops.Config != "" {
		config, err := LoadConfig(ops.Config)
		if err != nil {
			log.Print("[ERROR] " + fmt.Sprintf("Error loading CLI configuration: \n%s", err))
			return 1
		}
		c.Merge(config)
	}

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(c.Loglevel),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	Run()
	return 0
}
