package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if !strings.Contains(haystack, needle) {
		return -1
	} else {
		return strings.Index(haystack, needle)
	}
}
