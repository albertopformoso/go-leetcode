package main

import (
	"fmt"
	"sort"
)

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}

func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	l := len(nums)

	summation := func(nums []int, begin, end, target int) {
		for begin < end {
			if nums[begin]+nums[end]+target == 0 {
				r := make([]int, 0)
				r = append(r, nums[begin], nums[end], target)
				res = append(res, r)

				for begin < end && nums[begin] == nums[begin+1] {
					begin++
				}
				for begin < end && nums[end] == nums[end-1] {
					end--
				}
				begin++
				end--
			} else if target+nums[begin]+nums[end] < 0 {
				begin++
			} else {
				end--
			}
		}
	}

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		summation(nums, i+1, l-1, nums[i])
	}
	return
}
