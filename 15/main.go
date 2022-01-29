package main

import (
	"fmt"
	"sort"
)

func main() {
	var A = []int{-1, 2, -1, 0, 0, 3, 4, -1, -2, 1, 1000000}
	fmt.Println(threeSum(A))
}

func threeSum(nums []int) [][]int {
	//----i----L---R------len(nums)
	var result [][]int
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for {
			if L >= R {
				break
			}
			sum := nums[i] + nums[L] + nums[R]
			if sum > 0 {
				R--
				continue
			} else if sum < 0 {
				L++
				continue
			} else {
				result = append(result, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
			}
			L++
			R--
		}
	}
	return result
}
