[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/003-longest-substring-without-repeating-characters/003-longest-substring-without-repeating-characters.md) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/005-longest-palindromic-substring/005-longest-palindromic-substring.md)

# Median of Two Sorted Arrays

### Problem

https://leetcode.com/problems/median-of-two-sorted-arrays/

### Solution
```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var mergedArr []int = append(nums1, nums2...) 
	sort.Ints(mergedArr)

	index := len(mergedArr) / 2
	if len(mergedArr) % 2 == 1 {
		return float64(mergedArr[index])
	}

	return float64(mergedArr[index] + mergedArr[index - 1]) / 2
}
```