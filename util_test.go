package main

import (
	"reflect"
	"testing"
)

func TestIsFileExist(t *testing.T) {
	res := IsFileExist("/tmp/__foobarbaz")
	if res != false {
		t.Errorf("expected %s to eq %s", res, false)
	}
}

func TestCombineKeyValues(t *testing.T) {
	keys := []string{"a", "b", "c"}
	vals := []string{"A", "B", "C"}
	res := CombineKeyValues(keys, vals)
	expected := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %s to eq %s", res, expected)
	}
}

func TestReverseStrings(t *testing.T) {
	src := []string{"a", "b", "c"}
	expected := []string{"c", "b", "a"}
	res := ReverseStrings(src)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %s to eq %s", res, expected)
	}
}

func TestCopyMap(t *testing.T) {
	origin := map[string][]string{
		"A": []string{"a1", "b1", "c1"},
		"B": []string{"a2", "b2", "c2"},
	}
	copied := CopyMap(origin)
	copied["C"] = []string{"a3", "b3", "c3"}

	if len(copied) == len(origin) {
		t.Errorf("copied: %s, origin: %s", copied, origin)
	}
}
