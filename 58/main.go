package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	strs := strings.Split(strings.TrimRight(s, " "), " ")
	return len(strs[len(strs)-1])
}
