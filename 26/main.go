package main

import "fmt"

func main() {
	nums := []int{0, 0, 0, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	slow := 1
	for fast := 1; fast < length; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
