package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	num := []int{0, 2, 1, -3}
	target := 1
	fmt.Println(threeSumClosest(num, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	best := math.MaxInt32

	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L, R := i+1, len(nums)-1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				Rt := R - 1
				for L < Rt && nums[Rt] == nums[R] {
					Rt--
				}
				R = Rt
			} else {
				Lt := L + 1
				for Lt < R && nums[Lt] == nums[L] {
					Lt++
				}
				L = Lt
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
