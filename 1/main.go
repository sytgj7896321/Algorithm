package main

import (
	"fmt"
)

func main() {
	var A = []int{3, 3}
	var B []int
	var target = 6
	fmt.Println(twoSum2(A, target))
	fmt.Println(twoSum2(B, target))
}

//Tiny Memory
func twoSum(nums []int, target int) []int {
	var tag int
	for i := 0; i < len(nums); i++ {
		tag = i + 1
		for {
			if tag < len(nums) && nums[i]+nums[tag] != target {
				tag++
			} else if tag < len(nums) && nums[i]+nums[tag] == target {
				return []int{i, tag}
			} else {
				break
			}
		}
	}
	return nil
}

//Balance
func twoSum2(nums []int, target int) []int {
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
