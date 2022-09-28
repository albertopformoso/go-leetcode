[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/010-regular-expression-matching/010-regular-expression-matching.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/012-integer-to-roman/012-integer-to-roman.md)

# Container With Most Water

### Problem

https://leetcode.com/problems/container-with-most-water/

### Solution

```go
func maxArea(height []int) (maxArea int) {
	var left, right int = 0, len(height) - 1

	for left < right {
		base := right - left
		area := base * int(math.Min(float64(height[left]), float64(height[right])))
		maxArea = int(math.Max(float64(area), float64(maxArea)))

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return
}
```
