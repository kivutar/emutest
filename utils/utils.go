// Package utils contains utility functions that are used everywhere in the app.
package utils

// IndexOfString returns the index of a string in a string slice
func IndexOfString(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return 0
}
