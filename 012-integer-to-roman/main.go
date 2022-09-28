package main

import (
	"fmt"
)

func main() {
	res := intToRoman(32)
	fmt.Println(res)
}

func intToRoman(num int) (out string) {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i < len(roman); i++ {
		for num >= values[i] {
			out += roman[i]
			num -= values[i]
		}
	}
	return
}
