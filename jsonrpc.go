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
	req    *http.Request
	rpcReq jsonRPCRequest
}

// jsonRPCRequest is structure
type jsonRPCRequest struct {
	Jsonrpc string
	Method  string
	Params  interface{}
	ID      int
}

// ResponseFile returns file path
func (j *JSONRPC) ResponseFile() (string, map[string]string) {
	decoder := json.NewDecoder(j.req.Body)

	err := decoder.Decode(&j.rpcReq)
	if err != nil {
		log.Print("[ERROR] " + fmt.Sprintf("Error json decode: \n%s", err))
	}

	dict := j.dictionary()
	files := j.nominatedFiles(dict)

	for _, file := range files {
		if IsFileExist(file) {
			return file, dict
		}
	}

	return "", dict
}

// fileExt returns a extension
func (j *JSONRPC) fileExt() string {
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

	return ext
}

// dictionary returns resources
func (j *JSONRPC) dictionary() map[string]string {
	dict := make(map[string]string)
	var s string

	for _, v := range Config().Namespaces {
		err := scan.ScanTree(j.rpcReq.Params, "/"+v, &s)
		if err != nil {
			continue
		}
		dict[v] = s
	}

	return dict
}

// nominatedfiles returns file paths
func (j *JSONRPC) nominatedFiles(dict map[string]string) []string {
	var pathsOrderReal, pathsOrderVirt []string
	var src string

	ext := j.fileExt()
	keys := MapKeys(dict)
	vals := MapVals(dict)
	count := len(keys)

	src = path.Join(Config().Root, j.req.RequestURI, j.rpcReq.Method+ext)
	pathsOrderVirt = append(pathsOrderVirt, src)

	for i := 0; i <= count; i++ {
		normalPath := keys[:(count - i)]
		virtPath := vals[(count - i):]
		dir := strings.Join(append(normalPath, virtPath...), "/")
		src = path.Join(Config().Root, j.req.RequestURI, dir, j.rpcReq.Method+ext)
		pathsOrderVirt = append(pathsOrderVirt, src)
	}

	pathsOrderReal = ReverseStrings(pathsOrderVirt)
	for _, p := range pathsOrderReal {
		log.Print("[DEBUG] " + fmt.Sprintf("nominated file: %s", p))
	}

	return pathsOrderReal
}
