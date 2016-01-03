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
func (j *JSONRPC) ResponseFile(w http.ResponseWriter, r *http.Request) (string, string) {
	var file, dir string

	c := Config()

	decoder := json.NewDecoder(r.Body)
	var rpcReq jsonRPCRequest
	err := decoder.Decode(&rpcReq)
	if err != nil {
		fmt.Sprintf("Error json decode: \n\n%s", err)
	}

	for _, v := range c.Namespaces {
		if val, ok := rpcReq.Params[v]; ok {
			dir = path.Join(dir, val)
		}
	}

	file = path.Join(c.Root, r.RequestURI, dir, rpcReq.Method)
	if IsFileExist(file) {
		return file, ""
	}

	f, id := j.splitID(dir, c.AnonymousID)
	file = path.Join(c.Root, r.RequestURI, f, rpcReq.Method)
	return file, id
}

func (j *JSONRPC) splitID(file string, holder string) (string, string) {
	_, f := path.Split(path.Clean(file))
	d := path.Join(path.Dir(path.Clean(file)), holder)
	return d, f
}
