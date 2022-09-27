package main

import (
	"fmt"
)

func main() {
	res := isPalindrome(121)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	var reversed int
	var original int = x

	if x < 0 { return false }

	for x > 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}

	return original == reversed
}
