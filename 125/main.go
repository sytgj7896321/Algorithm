package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(test))

}

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	L := 0
	R := len(s) - 1
	for L < R {
		if isRight(s[L]) {
			if isRight(s[R]) {
				if s[L] == s[R] {
					L++
					R--
					continue
				} else {
					return false
				}
			} else {
				R--
				continue
			}
		} else {
			L++
			continue
		}
	}
	return true
}

func isRight(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
