package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"aaaa", "aaaaa", "aaaggs"}
	fmt.Println(longestCommonPrefix(a))
}

func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	minPos := 0
	for i, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
			minPos = i
		}
	}
	for _, str := range strs {
		for minLen > 0 {
			if strings.HasPrefix(str, strs[minPos][:minLen]) {
				break
			} else {
				minLen--
			}
		}
		if minLen == 0 {
			return ""
		}
	}
	return strs[minPos][:minLen]
}
