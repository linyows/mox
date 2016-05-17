package main

import "testing"

//func TestResponseFile(t *testing.T) {
//}
//
//func TestResponseExt(t *testing.T) {
//}
//
//func TestDictionary(t *testing.T) {
//}
//
//func TestNominatedFiles(t *testing.T) {
//}

func TestRemoveExt(t *testing.T) {
	rest := &REST{}
	file := "/foo/bar/baz.json"
	expected := "/foo/bar/baz"
	actual := rest.removeExt(file)
	if actual != expected {
		t.Errorf("expected %s to eq %s", actual, expected)
	}
}

//func localFormat(t *testing.T) {
//	rest := &REST{
//		req: *http.Request,
//	}
//	file := "baz.json"
//	expected := "GET--baz.json"
//	actual := rest.removeExt(file)
//	if actual != expected {
//		t.Errorf("expected %s to eq %s", actual, expected)
//	}
//}
