package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

// REST is structure
type REST struct {
}

// ResponseFile returns file path
func (re *REST) ResponseFile(w http.ResponseWriter, r *http.Request) (string, map[string]string) {
	c := Config()
	dict := make(map[string]string)
	params := make(map[string]string)
	var pathsOrderReal, pathsOrderVirt []string

	file := path.Join(c.Root, r.RequestURI+"--"+r.Method)
	if IsFileExist(file) {
		return file, dict
	}

	trimedURL := strings.TrimPrefix(r.RequestURI, "/")
	arrayPath := strings.Split(trimedURL, "/")
	for i, v := range arrayPath {
		if i%2 == 0 {
			continue
		}
		params[v] = arrayPath[i+1]
	}

	for _, v := range c.Namespaces {
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
		p := path.Join(c.Root, dir+"--"+r.Method)
		pathsOrderVirt = append(pathsOrderVirt, p)
	}

	pathsOrderReal = ReverseStrings(pathsOrderVirt)
	for _, p := range pathsOrderReal {
		log.Print("[DEBUG] " + fmt.Sprintf("search file: %s", p))
	}

	return file, dict
}
