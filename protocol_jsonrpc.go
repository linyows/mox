package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
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
func (j *JSONRPC) ResponseFile(w http.ResponseWriter, r *http.Request) (string, map[string]string) {
	c := Config()

	decoder := json.NewDecoder(r.Body)
	var rpcReq jsonRPCRequest
	err := decoder.Decode(&rpcReq)
	if err != nil {
		log.Print("[ERROR] " + fmt.Sprintf("Error json decode: \n%s", err))
	}

	var pathsOrderReal, pathsOrderVirt []string
	dict := make(map[string]string)

	for _, v := range c.Namespaces {
		if val, ok := rpcReq.Params[v]; ok {
			dict[v] = val
		}
	}

	keys := MapKeys(dict)
	vals := MapVals(dict)
	count := len(keys)

	for i := 0; i <= count; i++ {
		normalPath := keys[:(count - i)]
		virtPath := vals[(count - i):]
		dir := strings.Join(append(normalPath, virtPath...), "/")
		p := path.Join(c.Root, r.RequestURI, dir, rpcReq.Method)
		pathsOrderVirt = append(pathsOrderVirt, p)
	}

	pathsOrderReal = ReverseStrings(pathsOrderVirt)
	for _, p := range pathsOrderReal {
		log.Print("[DEBUG] " + fmt.Sprintf("search file: %s", p))
	}

	for _, file := range pathsOrderReal {
		if IsFileExist(file) {
			return file, dict
		}
	}

	return "", dict
}
