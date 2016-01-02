package main

import "fmt"

// Pox starts server.
func Pox(ops Ops) int {
	c := DefaultConfig()
	c.Set(ops)

	if ops.Config != "" {
		config, err := LoadConfig(ops.Config)
		if err != nil {
			fmt.Sprintf("Error loading CLI configuration: \n\n%s", err)
			return 1
		}
		c.Merge(config)
	}

	Run()

	return 0
}
