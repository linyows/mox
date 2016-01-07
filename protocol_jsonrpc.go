package main

import (
	"encoding/json"
	"fmt"
	"log"
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
func (j *JSONRPC) ResponseFile(w http.ResponseWriter, r *http.Request) (string, map[string]string) {
	c := Config()

	decoder := json.NewDecoder(r.Body)
	var rpcReq jsonRPCRequest
	err := decoder.Decode(&rpcReq)
	if err != nil {
		log.Print("[ERROR] " + fmt.Sprintf("Error json decode: \n%s", err))
	}

	pathsMap := make(map[string]string)
	dict := make(map[string]string)

	for _, v := range c.Namespaces {
		if val, ok := rpcReq.Params[v]; ok {
			dict[v] = val
		}
	}

	keysMap := dict.Keys()
	valsMap := dict.Vals()
	count := len(keysMap)

	for i, v := range keysMap {
		normalPath := keysMap[:(count - i)]
		virtPath := valsMap[(count - i):]
		dir := strings.Join(append(normalPath, virtPath...), "/")
		p := path.Join(c.Root, r.RequestURI, dir, rpcReq.Method)
		pathsMap = append(pathsMap, p)
	}

	for _, file := range pathsMap {
		if IsFileExist(file) {
			return file, dict
		}
	}

	return "", dict
}
