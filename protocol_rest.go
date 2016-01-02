package main

import (
	"net/http"
	"path"
)

// REST is structure
type REST struct {
}

// ResponseFile returns file path
func (re *REST) ResponseFile(w http.ResponseWriter, r *http.Request) string {
	return path.Join(Config().Root, r.RequestURI+"--"+r.Method)
}
