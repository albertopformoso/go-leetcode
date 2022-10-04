package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}[({})]"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var closers []rune
	open := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, v := range s {
		if closer, ok := open[v]; ok {
			closers = append(closers, closer)
			continue
		}

		l := len(closers) - 1
		if l < 0 || v != closers[l] {
			return false
		}
		
		closers = closers[:l]
	}

	return len(closers) == 0
}
