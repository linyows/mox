package main

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

type Config struct {
	Root  string
	Port  int
	Type  string
	Delay int
}

// LoadConfig loads the CLI configuration from "pox.conf" files.
func LoadConfig(path string) (*Config, error) {
	// Read the HCL file and prepare for parsing
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"Error reading %s: %s", path, err)
	}

	// Parse it
	obj, err := hcl.Parse(string(d))
	if err != nil {
		return nil, fmt.Errorf(
			"Error parsing %s: %s", path, err)
	}

	// Build up the result
	var result Config
	if err := hcl.DecodeObject(&result, obj); err != nil {
		return nil, err
	}

	return &result, nil
}
