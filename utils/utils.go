// Package utils contains utility functions that are used everywhere in the app.
package utils

import (
	"path/filepath"
)

// StringInSlice check wether a string is contain in a string slice.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IndexOfString returns the index of a string in a string slice
func IndexOfString(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return 0
}

// FileName returns the name of a file, without the path and extension.
func FileName(path string) string {
	name := filepath.Base(path)
	ext := filepath.Ext(name)
	name = name[0 : len(name)-len(ext)]
	return name
}
