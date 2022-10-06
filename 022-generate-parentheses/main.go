package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var out []string
	generator(&out, "", n, n)

	return out
}

func generator(out *[]string, s string, left, right int) {
	if left < 0 || right < 0 || left > right {
		return
	}

	if left == 0 && right == 0 {
		*out = append(*out, s)
		return
	}

	generator(out, s+"(", left-1, right)
	generator(out, s+")", left, right-1)
}
