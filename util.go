package main

import (
	"os"
)

// IsFileExist returns true if it is exists
func IsFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}

// CombineKeyValues returns Map
func CombineKeyValues(keys []string, values []string) map[string]string {
	keyValues := make(map[string]string)

	for i, key := range keys {
		keyValues[key] = values[i]
	}

	return keyValues
}

// ReverseStrings reverses array
func ReverseStrings(src []string) []string {
	var dst []string

	for i := len(src) - 1; i >= 0; i-- {
		dst = append(dst, src[i])
	}

	return dst
}
