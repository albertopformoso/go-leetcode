[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/009-palindrome-number/009-palindrome-number.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/011-container-with-most-water/011-container-with-most-water.md)

# Regular Expression Matching

### Problem

https://leetcode.com/problems/regular-expression-matching/

### Solution

```go
func isMatch(s, p string) bool {
	re := regexp.MustCompile(p)
	match := re.Find([]byte(s))
	return string(match) == s
}
```
