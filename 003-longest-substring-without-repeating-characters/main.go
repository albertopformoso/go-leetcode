package main

import (
	"fmt"
	"math"
)

func main() {
	res := lengthOfLongestSubstring("dvdf")
	fmt.Println(res)
}

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

