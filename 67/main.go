package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "11111"
	b := "101011111"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	sum := ""
	var lenMax, carry int
	lenA, lenB := len(a), len(b)
	if len(a) > len(b) {
		lenMax = lenA
	} else {
		lenMax = lenB
	}
	for i := 0; i < lenMax; i++ {
		if i < lenA {
			carry += int(a[lenA-1-i] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-1-i] - '0')
		}
		sum = strconv.Itoa(carry%2) + sum
		carry /= 2
	}
	if carry > 0 {
		sum = "1" + sum
	}
	return sum
}
