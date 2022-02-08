package main

import (
	"fmt"
)

func main() {
	digits := []int{1, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 > 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			break
		}
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
