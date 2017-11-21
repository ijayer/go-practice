package main

import "strings"

func main() {
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.HasSuffix(b, a) {
			return true
		}
	}
	return false
}
