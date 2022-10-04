package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
}

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