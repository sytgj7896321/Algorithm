package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 4
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}
