package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

// JSONRPC is structure
type JSONRPC struct {
}

// jsonRPCRequest is structure
type jsonRPCRequest struct {
	Jsonrpc string
	Method  string
	Params  map[string]string
	ID      int
}

// ResponseFile returns file path
func (j *JSONRPC) ResponseFile(w http.ResponseWriter, r *http.Request) string {
	c := Config()
	var p string

	decoder := json.NewDecoder(r.Body)
	var rpcReq jsonRPCRequest
	err := decoder.Decode(&rpcReq)
	if err != nil {
		fmt.Sprintf("Error json decode: \n\n%s", err)
	}

	for _, v := range c.Namespaces {
		if val, ok := rpcReq.Params[v]; ok {
			p = path.Join(p, val)
		}
	}

	return path.Join(c.Root, r.RequestURI, p, rpcReq.Method)
}
