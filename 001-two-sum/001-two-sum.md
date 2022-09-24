[Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/002-add-two-numbers)

# Two Sum

### Problem

https://leetcode.com/problems/two-sum/

### Solution
```go
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
```