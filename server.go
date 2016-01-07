package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// ResponseBody is structure
type ResponseBody struct {
	ID string
}

// This method handles all requests.
func handle(w http.ResponseWriter, r *http.Request) {
	var proto Protocol
	c := Config()

	log.Print("[INFO] " + fmt.Sprintf("%s - \"%s %s %s\" - \"%s\"",
		r.RemoteAddr, r.Method, r.RequestURI, r.Proto, strings.Join(r.Header["User-Agent"], ",")))
	log.Print("[DEBUG] " + fmt.Sprintf("%#v", c))

	if c.Delay > 0 {
		log.Print("[DEBUG] " + fmt.Sprintf("sleep %vs ...", c.Delay))
		time.Sleep(time.Duration(c.Delay) * time.Second)
	}

	switch c.Protocol {
	case "JSON-RPC":
		proto = &JSONRPC{}
	case "REST":
		proto = &REST{}
	default:
		panic(fmt.Sprintf("Error known protocol: %s", c.Protocol))
	}

	file, dict := proto.ResponseFile(w, r)

	log.Print("[DEBUG] " + fmt.Sprintf("file: %s", file))
	log.Print("[DEBUG] " + fmt.Sprintf("dict: %s", dict))

	ext := filepath.Ext(file)
	t := "Content-Type"

	for k, v := range c.Header {
		if k == t && ext != "" {
			w.Header().Set(t, mime.TypeByExtension(ext))
		} else {
			w.Header().Set(k, v)
		}
	}

	if file != "" && len(dict) > 0 {
		tpl := template.Must(template.ParseFiles(file))
		tpl.Execute(w, dict)
	} else if file == "" {
		fmt.Fprint(w, "custom 404")
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
