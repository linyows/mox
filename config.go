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

// Merge merges other configurations it self.
func (c *Config) Merge(otherConfig *Config) *Config {
	if otherConfig.Root != "" {
		c.Root = otherConfig.Root
	}
	if otherConfig.Addr != "" {
		c.Addr = otherConfig.Addr
	}
	if otherConfig.Type != "" {
		c.Type = otherConfig.Type
	}
	if otherConfig.Delay != 0 {
		c.Delay = otherConfig.Delay
	}
	if otherConfig.Loglevel != "" {
		c.Loglevel = otherConfig.Loglevel
	}

	return c
}

// Set sets from Ops
func (c *Config) Set(o Ops) *Config {
	if o.Root != "" {
		c.Root = o.Root
	}
	if o.Addr != "" {
		c.Addr = o.Addr
	}
	if o.Type != "" {
		c.Type = o.Type
	}
	if o.Delay != 0 {
		c.Delay = o.Delay
	}
	if o.Loglevel != "" {
		c.Loglevel = o.Loglevel
	}

	return c
}
