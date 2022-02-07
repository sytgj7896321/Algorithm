package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 11111
	fmt.Println(isPalindrome(a))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] == str[length-1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}
