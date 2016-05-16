package main

// Protocol is interface
type Protocol interface {
	ResponseFile() (string, map[string]string)
	ResponseExt() string
}
