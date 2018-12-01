package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hashicorp/hcl"
)

var instance *config

// Config is the structure of the configuration for CLI.
type config struct {
	Root        string
	Addr        string
	Protocol    string
	Delay       int
	LogLevel    string
	AnonymousID string
	Header      map[string]string
	Namespaces  []string
}

// DefaultConfig returns default structure.
func DefaultConfig() *config {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	instance = &config{
		Root:        "/var/www/mox",
		Addr:        "localhost:8080",
		Protocol:    "REST",
		Delay:       0,
		LogLevel:    "INFO",
		AnonymousID: "ANONID",
		Header: map[string]string{
			"Server":       hostname,
			"Content-Type": "application/octet-stream",
			"X-Served-By":  "mox",
		},
	}

	return instance
}

// LoadConfig loads the CLI configuration from conf files.
func LoadConfig(path string) (*config, error) {
	// Read the HCL file and prepare for parsing
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading %s: %s", path, err)
	}

	// Parse it
	obj, err := hcl.Parse(string(d))
	if err != nil {
		return nil, fmt.Errorf("Error parsing %s: %s", path, err)
	}

	// Build up the result
	var result config
	if err := hcl.DecodeObject(&result, obj); err != nil {
		return nil, err
	}

	return &result, nil
}

// Merge merges other configurations it self.
func (c *config) Merge(otherConfig *config) *config {
	if otherConfig.Root != "" {
		c.Root = otherConfig.Root
	}
	if otherConfig.Addr != "" {
		c.Addr = otherConfig.Addr
	}
	if otherConfig.Protocol != "" {
		c.Protocol = strings.ToUpper(otherConfig.Protocol)
	}
	if otherConfig.Delay != 0 {
		c.Delay = otherConfig.Delay
	}
	if otherConfig.LogLevel != "" {
		c.LogLevel = strings.ToUpper(otherConfig.LogLevel)
	}
	if otherConfig.AnonymousID != "" {
		c.AnonymousID = otherConfig.AnonymousID
	}
	if len(otherConfig.Namespaces) != 0 {
		c.Namespaces = otherConfig.Namespaces
	}
	for k, v := range otherConfig.Header {
		c.Header[k] = v
	}

	return c
}

// Set sets from Options
func (c *config) Set(o Options) *config {
	if o.Root != "" {
		c.Root = o.Root
	}
	if o.Addr != "" {
		c.Addr = o.Addr
	}
	if o.Protocol != "" {
		c.Protocol = strings.ToUpper(o.Protocol)
	}
	if o.Delay != 0 {
		c.Delay = o.Delay
	}
	if o.LogLevel != "" {
		c.LogLevel = strings.ToUpper(o.LogLevel)
	}

	return c
}

// SetFromEnv sets from env variables
func (c *config) SetFromEnv() *config {
	upperName := strings.ToUpper("mox")

	root := os.Getenv(upperName + "_ROOT")
	if root != "" {
		c.Root = root
	}

	addr := os.Getenv(upperName + "_ADDR")
	if addr != "" {
		c.Addr = addr
	}

	return c
}

// Config returns config singleton structure.
func Config() *config {
	if instance == nil {
		return DefaultConfig()
	}
	return instance
}
