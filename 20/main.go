package main

import "fmt"

func main() {
	a := "{}[]()"
	fmt.Println(isValid(a))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if m[s[i]] > 0 {
			if len(stack) == 0 || m[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
