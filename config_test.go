package main

import "testing"

func TestDefaultConfig(t *testing.T) {
	c := DefaultConfig()
	root := "/var/www/mox"
	if c.Root != root {
		t.Errorf("Root expected %s, got %s.", root, c.Root)
	}

	addr := "localhost:8080"
	if c.Addr != addr {
		t.Errorf("Addr expected %s, got %s.", addr, c.Addr)
	}

	protocol := "REST"
	if c.Protocol != protocol {
		t.Errorf("Protocol expected %s, got %s.", protocol, c.Protocol)
	}

	delay := 0
	if c.Delay != delay {
		t.Errorf("Delay expected %v, got %v.", delay, c.Delay)
	}

	logLevel := "INFO"
	if c.LogLevel != logLevel {
		t.Errorf("LogLevel expected %s, got %s.", logLevel, c.LogLevel)
	}

	anonymousID := "ANONID"
	if c.AnonymousID != anonymousID {
		t.Errorf("AnonymousID expected %s, got %s.", anonymousID, c.AnonymousID)
	}

	header := map[string]string{
		"Content-Type": "application/octet-stream",
		"X-Served-By":  "mox",
	}
	if c.Header["Content-Type"] != header["Content-Type"] {
		t.Errorf("Content-Type Header expected %s, got %s.", header["Content-Type"], c.Header["Content-Type"])
	}
	if c.Header["X-Served-By"] != header["X-Served-By"] {
		t.Errorf("X-Served-By Header expected %s, got %s.", header["X-Served-By"], c.Header["X-Served-By"])
	}
}

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig("testdata/rest/mox.conf")
	if err != nil {
		t.Errorf("Raised error %#v.", err)
	}
	server := "REST API"
	if c.Header["Server"] != server {
		t.Errorf("Server Header expected %s, got %s.", server, c.Header["Server"])
	}
}

func TestMerge(t *testing.T) {
	c := DefaultConfig()
	c2, _ := LoadConfig("testdata/rest/mox.conf")
	c.Merge(c2)
	if c.LogLevel != "DEBUG" {
		t.Errorf("It is not override config: expected %s, got %s.", "DEBUG", c.LogLevel)
	}
}

func TestSet(t *testing.T) {
	c := DefaultConfig()
	opt := Ops{
		Root: "/mox",
	}
	c.Set(opt)
	if c.Root != "/mox" {
		t.Errorf("It is not set for root: %s", c.Root)
	}
}

func TestSetFromEnv(t *testing.T) {
	c := DefaultConfig()
	c.SetFromEnv()
	if c.Root != "/var/www/mox" {
		t.Errorf("It is not set for root: %s", c.Root)
	}
}

func TestConfig(t *testing.T) {
	c1 := DefaultConfig()
	c2 := Config()
	if c1 != c2 {
		t.Errorf("It is not same.")
	}
	c2.Root = "/mox"
	c3 := Config()
	if c3.Root != "/mox" {
		t.Errorf("It is not singleton.")
	}
}
