package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

// JSONRPC is structure
type JSONRPC struct {
	req *http.Request
}

// jsonRPCRequest is structure
type jsonRPCRequest struct {
	Jsonrpc string
	Method  string
	Params  map[string]string
	ID      int
}

func (j *JSONRPC) targetFile(method string, dir string) string {
	var ext string
	var mimeType string

	reqURI := j.req.RequestURI
	clientMimes := j.req.Header["Content-Type"]
	serverMime := Config().Header["Content-Type"]
	extByURI := filepath.Ext(reqURI)

	if extByURI == "" {
		if len(clientMimes) != 0 {
			mimeType = clientMimes[0]
		} else {
			mimeType = serverMime
		}

		exts, err := mime.ExtensionsByType(mimeType)
		if err != nil {
			fmt.Errorf("Error mime to ext %s: %s", mimeType, err)
		} else {
			ext = exts[0]
		}
	} else {
		ext = extByURI
	}

	return path.Join(Config().Root, j.req.RequestURI, dir, method+ext)
}

// ResponseFile returns file path
func (j *JSONRPC) ResponseFile() (string, map[string]string) {
	decoder := json.NewDecoder(j.req.Body)
	var rpcReq jsonRPCRequest
	err := decoder.Decode(&rpcReq)
	if err != nil {
		log.Print("[ERROR] " + fmt.Sprintf("Error json decode: \n%s", err))
	}

	var pathsOrderReal, pathsOrderVirt []string
	dict := make(map[string]string)

	for _, v := range Config().Namespaces {
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
		pathsOrderVirt = append(pathsOrderVirt, j.targetFile(rpcReq.Method, dir))
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
