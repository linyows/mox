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

	"github.com/mattn/go-scan"
)

// JSONRPC is structure
type JSONRPC struct {
	req         *http.Request
	rpcReq      jsonRPCRequest
	responseExt string
}

// jsonRPCRequest is structure
type jsonRPCRequest struct {
	Jsonrpc string
	Method  string
	Params  interface{}
	ID      int
}

// ResponseFile returns file path
func (j *JSONRPC) ResponseFile() (string, map[string][]string) {
	decoder := json.NewDecoder(j.req.Body)

	err := decoder.Decode(&j.rpcReq)
	if err != nil {
		log.Print("[ERROR] " + fmt.Sprintf("Error json decode: \n%s", err))
	}

	dict := j.dictionary()
	files := j.nominatedFiles(CopyMap(dict))

	for _, file := range files {
		if IsFileExist(file) {
			return file, dict
		}
	}

	return "", dict
}

// ResponseExt returns a extension
func (j *JSONRPC) ResponseExt() string {
	if j.responseExt != "" {
		return j.responseExt
	}

	ext := filepath.Ext(j.req.RequestURI)

	if ext == "" {
		mimeType := Config().Header["Content-Type"]

		exts, err := mime.ExtensionsByType(mimeType)
		if err != nil {
			log.Print("[ERROR] " + fmt.Sprintf("Error mime to ext %s: \n%s", mimeType, err))
		} else {
			j.responseExt = exts[0]
		}
	} else {
		j.responseExt = ext
	}

	return j.responseExt
}

// dictionary returns resources
func (j *JSONRPC) dictionary() map[string][]string {
	var keys []string
	var vals []string
	var s string

	for _, v := range Config().Namespaces {
		err := scan.ScanTree(j.rpcReq.Params, "/"+v, &s)
		if err != nil {
			log.Print("[DEBUG] " + fmt.Sprintf("JSON parsed, but \"%s\" not found -- %s", v, err))
			continue
		}
		keys = append(keys, v)
		vals = append(vals, s)
	}

	return map[string][]string{
		"keys":   keys,
		"values": vals,
	}
}

// nominatedfiles returns file paths
func (j *JSONRPC) nominatedFiles(dict map[string][]string) []string {
	var pathsOrderReal, pathsOrderVirt []string
	var src string

	keys := dict["keys"]
	vals := dict["values"]
	count := len(dict["keys"])

	src = path.Join(Config().Root, j.req.RequestURI, j.rpcReq.Method+j.ResponseExt())
	pathsOrderVirt = append(pathsOrderVirt, src)

	for i := 0; i <= count; i++ {
		normalPath := keys[:(count - i)]
		virtPath := vals[(count - i):]
		dir := strings.Join(append(normalPath, virtPath...), "/")
		src = path.Join(Config().Root, j.req.RequestURI, dir, j.rpcReq.Method+j.ResponseExt())
		pathsOrderVirt = append(pathsOrderVirt, src)
	}

	pathsOrderReal = ReverseStrings(pathsOrderVirt)
	for _, p := range pathsOrderReal {
		log.Print("[DEBUG] " + fmt.Sprintf("Nominated file: %s", p))
	}

	return pathsOrderReal
}
