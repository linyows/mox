package main

import "fmt"

func Pox(ops Ops) int {
	if ops.Config == "" {
		ops.Config = "~/.go/src/github.com/linyows/pox/examples/pox.conf"
	}

	if ops.Config != "" {
		config, err := LoadConfig(ops.Config)
		if err != nil {
			fmt.Sprintf("Error loading CLI configuration: \n\n%s", err)
			return 1
		}
		fmt.Printf("%+v", config)
	}

	s := NewServer()
	s.Start()

	return 0
}
