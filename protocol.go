package main

import "net/http"

// Protocol is interface
type Protocol interface {
	ResponseFile(w http.ResponseWriter, r *http.Request) (string, map[string]string)
}
