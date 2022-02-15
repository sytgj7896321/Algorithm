package main

import "fmt"

func main() {
	test := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(test))
}

func maxArea(height []int) int {
	if len(height) == 2 {
		return findSmall(height[0], height[1])
	}
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		small := findSmall(height[l], height[r])
		now := (r - l) * small
		if now > max {
			max = now
		}
		if small == height[l] {
			l++
		} else {
			r--
		}
	}
	return max
}

func findSmall(l, r int) int {
	if l <= r {
		return l
	} else {
		return r
	}
}
