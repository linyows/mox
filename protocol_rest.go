package main

import (
	"net/http"
	"path"
)

// REST is structure
type REST struct {
}

// ResponseFile returns file path
func (re *REST) ResponseFile(w http.ResponseWriter, r *http.Request) (string, string) {
	file := path.Join(Config().Root, r.RequestURI+"--"+r.Method)

	c := Config()

	if IsFileExist(file) {
		return file, ""
	}

	f, id := re.splitID(r.RequestURI, c.AnonymousID)
	file = path.Join(c.Root, f+"--"+r.Method)
	return file, id
}

func (re *REST) splitID(file string, holder string) (string, string) {
	_, f := path.Split(path.Clean(file))
	d := path.Join(path.Dir(path.Clean(file)), holder)
	return d, f
}
