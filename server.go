package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
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

	log.Print("[INFO] " + fmt.Sprintf("%s - \"%s %s %s\" - \"%s\"",
		r.RemoteAddr, r.Method, r.RequestURI, r.Proto, strings.Join(r.Header["User-Agent"], ",")))
	log.Print("[DEBUG] " + fmt.Sprintf("%#v", Config()))

	if Config().Delay > 0 {
		log.Print("[DEBUG] " + fmt.Sprintf("sleep %vs ...", Config().Delay))
		time.Sleep(time.Duration(Config().Delay) * time.Second)
	}

	switch Config().Protocol {
	case "JSON-RPC":
		proto = &JSONRPC{
			req: r,
		}

	case "REST":
		proto = &REST{
			req: r,
		}
	default:
		panic(fmt.Sprintf("Error known protocol: %s", Config().Protocol))
	}

	file, dict := proto.ResponseFile()

	log.Print("[DEBUG] " + fmt.Sprintf("file: %s", file))
	log.Print("[DEBUG] " + fmt.Sprintf("dict: %s", dict))

	t := "Content-Type"
	for k, v := range Config().Header {
		if k == t {
			w.Header().Set(t, mime.TypeByExtension(proto.ResponseExt()))
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
	s := &http.Server{
		Addr:           Config().Addr,
		Handler:        http.HandlerFunc(handle),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
