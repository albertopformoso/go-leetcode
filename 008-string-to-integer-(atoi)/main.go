package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	res := myAtoi("9223372036854775808")
	fmt.Printf("num: %v | Type: %[1]T\n", res)
}

func myAtoi(s string) (out int) {
	var i int
	var sign int = 1

	s = strings.TrimSpace(s)

	if len(s) == 0 { return out }
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for ; i<len(s); i++ {
		if sign*out >= math.MaxInt32 { return math.MaxInt32 }
		if sign*out <= math.MinInt32 { return math.MinInt32 }

		if s[i] < '0' || s[i] > '9' { break }

		out = out*10 + (int(s[i]) - '0')
	}

	if sign*out >= math.MaxInt32 { return math.MaxInt32 }
	if sign*out <= math.MinInt32 { return math.MinInt32 }

	return out * sign
}
