package main

import (
	"fmt"
	"sort"
)

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
	fmt.Println(num1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(append(nums1[:m], nums2[:n]...))
}
