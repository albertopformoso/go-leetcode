package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	res := threeSumClosest([]int{-1, 2, 1, 4}, 1)
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {
	var out int
	dist := math.MaxInt32
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		begin, end := i+1, l-1
		for begin < end {
			sum := nums[begin] + nums[end] + nums[i]
			if sum < target {
				if target-sum < dist {
					dist = target - sum
					out = sum
				}
				begin++
			} else if sum > target {
				if sum-target < dist {
					dist = sum - target
					out = sum
				}
				end--
			} else {
				return sum
			}
		}
	}

	return out
}
