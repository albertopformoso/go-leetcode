package main

import (
	"fmt"
	"math"
)

func main() {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}

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
