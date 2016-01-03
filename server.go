package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

// ResponseBody is structure
type ResponseBody struct {
	ID string
}

// This method handles all requests.
func handle(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] " + fmt.Sprintf("%#v", r))

	var proto Protocol
	c := Config()

	if c.Delay > 0 {
		time.Sleep(time.Duration(c.Delay) * time.Second)
	}

	for k, v := range c.Header {
		w.Header().Set(k, v)
	}

	switch c.Protocol {
	case "JSON-RPC":
		proto = &JSONRPC{}
	case "REST":
		proto = &REST{}
	default:
		panic(fmt.Sprintf("Error known protocol: %s", c.Protocol))
	}

	file, id := proto.ResponseFile(w, r)

	ext := filepath.Ext(file)
	if ext != "" {
		contentType := mime.TypeByExtension(ext)
		w.Header().Set("Content-Type", contentType)
	}

	if id != "" && IsFileExist(file) {
		tpl := template.Must(template.ParseFiles(file))
		tpl.Execute(w, ResponseBody{ID: id})
	} else {
		http.ServeFile(w, r, file)
	}
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
