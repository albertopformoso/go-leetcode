package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var mergedArr []int = append(nums1, nums2...) 
	sort.Ints(mergedArr)

	index := len(mergedArr) / 2
	if len(mergedArr) % 2 == 1 {
		return float64(mergedArr[index])
	}

	return float64(mergedArr[index] + mergedArr[index - 1]) / 2
}

