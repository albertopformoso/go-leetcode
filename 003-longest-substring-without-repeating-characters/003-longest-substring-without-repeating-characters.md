[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/002-add-two-numbers/002-add-two-numbers.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/004-median-of-two-sorted-arrays/004-median-of-two-sorted-arrays.md)

# Add Two Numbers

### Problem

https://leetcode.com/problems/longest-substring-without-repeating-characters/

### Solution
```go
func lengthOfLongestSubstring(s string) int {
	var count float64
	var aux = make(map[rune]int)

	var end int
	for i, c := range s {
		val, found := aux[c]
		aux[c] = i

		if !found {
			end++
			count = math.Max(count, float64(end))
			continue
		}

		if i - end <= val {
			end = i - val
		} else {
			end++
			count = math.Max(count, float64(end))
		}
	}

	return int(count)
}
```