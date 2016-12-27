package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/logutils"
)

// Mox starts server.
func Mox(opt Options) int {
	c := DefaultConfig()
	c.Set(opt)

	if opt.Config != "" {
		config, err := LoadConfig(opt.Config)
		if err != nil {
			log.Print("[ERROR] " + fmt.Sprintf("Error loading CLI configuration: \n%s", err))
			return 1
		}
		c.Merge(config)
	}

	c.SetFromEnv()

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(c.LogLevel),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	Run()
	return 0
}
