package main

import (
	"os"
)

// IsFileExist returns true if it is exists
func IsFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}
