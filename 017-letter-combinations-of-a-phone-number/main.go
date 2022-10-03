package main

import (
	"fmt"
)

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}

var letters = []string{0: "", 1: "", 2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz"}

func letterCombinations(digits string) (out []string) {

	if len(digits) > 0 {
		helper(&out, digits, "", 0)
	}

	return
}

func helper(out *[]string, digits, curr string, index int) {
	if index == len(digits) {
		*out = append(*out, curr)
		return
	}

	temp := letters[digits[index]-48]
	for _, t := range temp {
		helper(out, digits, curr+string(t), index+1)
	}
}
