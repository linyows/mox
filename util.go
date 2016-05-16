package main

import (
	"os"
)

// IsFileExist returns true if it is exists
func IsFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}

// MapKeys returns array
func MapKeys(m map[string]string) []string {
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// MapVals returns array
func MapVals(m map[string]string) []string {
	var vals []string

	for _, v := range m {
		vals = append(vals, v)
	}

	return vals
}

// ReverseStrings reverses array
func ReverseStrings(src []string) []string {
	var dst []string

	for i := len(src) - 1; i >= 0; i-- {
		dst = append(dst, src[i])
	}

	return dst
}
