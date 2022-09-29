package main

import (
	"fmt"
)

func main() {
	res := romanToInt("III")
	fmt.Println(res)
}

func romanToInt(s string) (result int) {
	var romanInt = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	for i := 0; i < len(s); i++ {
		j := i + 1
		if j >= len(s) { j = i }

		if romanInt[s[i]] < romanInt[s[j]] {
			result += romanInt[s[j]] - romanInt[s[i]]
			i++
			continue
		}
		result += romanInt[s[i]]
	}

	return
}
