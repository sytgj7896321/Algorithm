package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IX"))
}

func romanToInt(s string) int {
	var sum, plus int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			plus = 1
			sum += plus
		case 'V':
			if i != 0 && s[i-1] == 'I' {
				plus = 3
				sum += plus
				continue
			}
			plus = 5
			sum += plus
		case 'X':
			if i != 0 && s[i-1] == 'I' {
				plus = 8
				sum += plus
				continue
			}
			plus = 10
			sum += plus
		case 'L':
			if i != 0 && s[i-1] == 'X' {
				plus = 30
				sum += plus
				continue
			}
			plus = 50
			sum += plus
		case 'C':
			if i != 0 && s[i-1] == 'X' {
				plus = 80
				sum += plus
				continue
			}
			plus = 100
			sum += plus
		case 'D':
			if i != 0 && s[i-1] == 'C' {
				plus = 300
				sum += plus
				continue
			}
			plus = 500
			sum += plus
		case 'M':
			if i != 0 && s[i-1] == 'C' {
				plus = 800
				sum += plus
				continue
			}
			plus = 1000
			sum += plus
		}
	}
	return sum
}
