package main

import (
	"fmt"
	"log"
	"mime"
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
	var mimeType string
	var URI string

	reqURI := re.req.RequestURI
	clientMimes := re.req.Header["Content-Type"]
	serverMime := Config().Header["Content-Type"]
	method := re.req.Method
	extByURI := filepath.Ext(reqURI)
	URI = strings.Trim(reqURI, "/")

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
		URI = URI + ext
	}

	dir, file := path.Split(URI)

	return path.Join(Config().Root, dir, method+"--"+file)
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

	trimedURL := strings.Trim(re.req.RequestURI, "/")
	arrayPath := strings.Split(trimedURL, "/")

	for i, v := range arrayPath {

		if i != 0 && i%2 != 0 {
			continue
		}

		next := i + 1
		if len(arrayPath) <= next {
			break
		}
		params[v] = arrayPath[next]
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
		ext := filepath.Ext(file)

		p := path.Join(Config().Root, dir, re.req.Method+"--"+Config().AnonymousID+ext)
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
