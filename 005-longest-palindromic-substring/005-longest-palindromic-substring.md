[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/004-median-of-two-sorted-arrays/004-median-of-two-sorted-arrays.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/006-zigzag-conversion/006-zigzag-conversion.md)

# Longest Palindromic Substring

### Problem

https://leetcode.com/problems/longest-palindromic-substring/

### Solution
```go
func longestPalindrome(s string) string {
	if s == "" || len(s) < 1 {
		return s
	}

	var start, end int
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i , i+1)
		len := int(math.Max(float64(len1), float64(len2)))
		if len > end - start {
			start = i - (len - 1) / 2
			end = i + len / 2
		}
	}

	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
```