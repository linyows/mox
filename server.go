package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// jsonRPCRequestBody is structure
type jsonRPCRequestBody struct {
	Jsonrpc string
	Method  string
	Params  map[string]string
	ID      int
}

// This method handles all requests.
func handle(w http.ResponseWriter, r *http.Request) {
	var p string
	c := Config()

	if c.Delay > 0 {
		time.Sleep(time.Duration(c.Delay) * time.Second)
	}

	if strings.ToUpper(c.Type) == "JSON-RPC" {
		decoder := json.NewDecoder(r.Body)
		var b jsonRPCRequestBody
		err := decoder.Decode(&b)
		if err != nil {
			fmt.Sprintf("Error json decode: \n\n%s", err)
		}

		for _, v := range c.Namespaces {
			if val, ok := b.Params[v]; ok {
				p = path.Join(p, val)
			}
		}

		p = path.Join(c.Root, r.RequestURI, p, b.Method)
	} else {
		p = path.Join(c.Root, r.RequestURI+"--"+r.Method)
	}

	ext := filepath.Ext(p)

	for k, v := range c.Header {
		w.Header().Set(k, v)
	}

	if ext != "" {
		contentType := mime.TypeByExtension(ext)
		w.Header().Set("Content-Type", contentType)
	}

	http.ServeFile(w, r, p)
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
