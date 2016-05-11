package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

// REST is structure
type REST struct {
	req *http.Request
}

func (re *REST) targetFile() string {
	var ext string
	ext = filepath.Ext(re.req.RequestURI)

	if ext == "" {
		if len(re.req.Header["Content-Type"]) != 0 {
			ext = re.req.Header["Content-Type"][0]
		} else {
			ext = Config().Header["Content-Type"]
		}
	}

	return path.Join(Config().Root, re.req.Method+"--"+re.req.RequestURI+ext)
}

// ResponseFile returns file path
func (re *REST) ResponseFile() (string, map[string]string) {
	dict := make(map[string]string)
	params := make(map[string]string)
	var pathsOrderReal, pathsOrderVirt []string

	file := re.targetFile()
	if IsFileExist(file) {
		return file, dict
	}

	trimedURL := strings.TrimPrefix(re.req.RequestURI, "/")
	arrayPath := strings.Split(trimedURL, "/")
	for i, v := range arrayPath {
		if i%2 == 0 {
			continue
		}
		params[v] = arrayPath[i+1]
	}

	for _, v := range Config().Namespaces {
		if val, ok := params[v]; ok {
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
		p := path.Join(Config().Root, dir+re.req.Method+"--")
		pathsOrderVirt = append(pathsOrderVirt, p)
	}

	pathsOrderReal = ReverseStrings(pathsOrderVirt)
	for _, p := range pathsOrderReal {
		log.Print("[DEBUG] " + fmt.Sprintf("search file: %s", p))
	}

	return file, dict
}
