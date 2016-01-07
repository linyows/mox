package main

import (
	"net/http"
	"path"
)

// REST is structure
type REST struct {
}

// ResponseFile returns file path
func (re *REST) ResponseFile(w http.ResponseWriter, r *http.Request) (string, map[string]string) {
	file := path.Join(Config().Root, r.RequestURI+"--"+r.Method)

	//c := Config()
	dict := make(map[string]string)

	//if IsFileExist(file) {
	return file, dict
	//}

	//f, id := re.splitID(r.RequestURI, c.AnonymousID)
	//file = path.Join(c.Root, f+"--"+r.Method)
	//return file, dict
}
