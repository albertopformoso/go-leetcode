[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/015-3sum/015-3sum.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/017-letter-combinations-of-a-phone-number/017-letter-combinations-of-a-phone-number.md)

# 3Sum Closest

### Problem

https://leetcode.com/problems/3sum-closest/

### Solution

```go
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
```
