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
	req         *http.Request
	responseExt string
}

// ResponseFile returns file path
func (re *REST) ResponseFile() (string, map[string][]string) {
	reqURI := strings.Trim(re.req.RequestURI, "/") + re.ResponseExt()
	dir, file := path.Split(reqURI)
	src := path.Join(Config().Root, dir, re.localFormat(file))

	if IsFileExist(src) {
		return src, make(map[string][]string)
	}

	dict := re.dictionary(reqURI)
	srcs := re.nominatedFiles(dict)

	for _, src := range srcs {
		if IsFileExist(src) {
			return src, dict
		}
	}

	return "", dict
}

// ResponseExt returns file extension
func (re *REST) ResponseExt() string {
	if re.responseExt != "" {
		return re.responseExt
	}

	var mimeType string
	clientMimes := re.req.Header["Content-Type"]
	extByURI := filepath.Ext(re.req.RequestURI)

	if extByURI == "" {
		if len(clientMimes) != 0 {
			mimeType = clientMimes[0]
		} else {
			mimeType = Config().Header["Content-Type"]
		}
		exts, err := mime.ExtensionsByType(mimeType)
		if err != nil {
			log.Print("[ERROR] " + fmt.Sprintf("Error mime to ext %s: \n%s", mimeType, err))
			return ""
		}
		re.responseExt = exts[0]
	} else {
		re.responseExt = extByURI
	}

	return re.responseExt
}

// dictionary returns resources
func (re *REST) dictionary(URI string) map[string][]string {
	var keys []string
	var vals []string
	params := make(map[string]string)

	arrayPath := strings.Split(URI, "/")

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
			keys = append(keys, v)
			vals = append(vals, re.removeExt(val))
		}
	}

	return map[string][]string{
		"keys":   keys,
		"values": vals,
	}
}

// nominatedfiles returns file paths
func (re *REST) nominatedFiles(dict map[string][]string) []string {
	var pathsOrderReal, pathsOrderVirt []string

	keys := dict["keys"]
	vals := dict["values"]
	count := len(dict["keys"])

	for i := 0; i <= count; i++ {
		normalPath := keys[:(count - i)]
		virtPath := vals[(count - i):]
		dir := strings.Join(append(normalPath, virtPath...), "/")
		path := path.Join(Config().Root, dir, re.localFormat(Config().AnonymousID+re.ResponseExt()))
		pathsOrderVirt = append(pathsOrderVirt, path)
	}

	pathsOrderReal = ReverseStrings(pathsOrderVirt)
	for _, p := range pathsOrderReal {
		log.Print("[DEBUG] " + fmt.Sprintf("Nominated file: %s", p))
	}

	return pathsOrderReal
}

// removeExt returns file name without extension
func (re *REST) removeExt(file string) string {
	return file[0 : len(file)-len(filepath.Ext(file))]
}

// localFormat returns path for local
func (re *REST) localFormat(file string) string {
	return re.req.Method + "--" + file
}
