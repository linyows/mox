package main

import (
	"bytes"
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

	log.Print("[DEBUG] " + fmt.Sprintf("File: %s", file))
	log.Print("[DEBUG] " + fmt.Sprintf("Dict: %s", dict))

	t := "Content-Type"
	for k, v := range Config().Header {
		if k == t {
			w.Header().Set(t, mime.TypeByExtension(proto.ResponseExt()))
		} else {
			w.Header().Set(k, v)
		}
	}

	if file != "" && len(dict["keys"]) > 0 {
		tpl := template.Must(template.ParseFiles(file))
		bufbody := new(bytes.Buffer)
		tpl.Execute(bufbody, CombineKeyValues(dict["keys"], dict["values"]))
		body := bufbody.String()
		log.Print("[DEBUG] " + fmt.Sprintf("Response: \n%s", body))
		fmt.Fprint(w, body)
	} else if file == "" {
		log.Print("[DEBUG] " + fmt.Sprintf("Response: %s", "404: Not Found"))
		fmt.Fprint(w, "404: Not Found")
	} else {
		http.ServeFile(w, r, file)
	}
}

// Run server
func Run() {
	log.Print("[DEBUG] " + fmt.Sprintf("%#v", Config()))

	s := &http.Server{
		Addr:           Config().Addr,
		Handler:        http.HandlerFunc(handle),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
