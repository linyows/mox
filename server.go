package main

import (
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"time"
)

// This method handles all requests.
func handle(w http.ResponseWriter, r *http.Request) {
	var protocol Protocol
	c := Config()

	if c.Delay > 0 {
		time.Sleep(time.Duration(c.Delay) * time.Second)
	}

	for k, v := range c.Header {
		w.Header().Set(k, v)
	}

	switch c.Protocol {
	case "JSON-RPC":
		protocol = &JSONRPC{}
	case "REST":
		protocol = &REST{}
	default:
		panic(fmt.Sprintf("Error known protocol: %s", c.Protocol))
	}

	file := protocol.ResponseFile(w, r)
	ext := filepath.Ext(file)

	if ext != "" {
		contentType := mime.TypeByExtension(ext)
		w.Header().Set("Content-Type", contentType)
	}

	http.ServeFile(w, r, file)
}

// Run server
func Run() {
	c := Config()

	s := &http.Server{
		Addr:           c.Addr,
		Handler:        http.HandlerFunc(handle),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
