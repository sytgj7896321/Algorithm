package main

import (
	"fmt"
)

func main() {
	var A = []int{3, 3}
	var B []int
	var target = 6
	fmt.Println(twoSum(A, target))
	fmt.Println(twoSum(B, target))
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, v := range nums {
		value, ok := dict[target-v]
		if ok && value != i {
			return []int{value, i}
		}
		dict[v] = i
	}
	return nil
}
