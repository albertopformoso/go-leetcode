package main

import (
	"fmt"
	"regexp"
)

func main() {
	res := isMatch("ab", ".*")
	fmt.Println(res)
}

func isMatch(s, p string) bool {
	re := regexp.MustCompile(p)
	match := re.Find([]byte(s))
	return string(match) == s
}
