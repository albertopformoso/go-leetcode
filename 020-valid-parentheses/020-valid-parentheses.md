[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/019-remove-nth-node-from-end-of-list/019-remove-nth-node-from-end-of-list.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next]()

# Valid Parentheses

### Problem

https://leetcode.com/problems/valid-parentheses/

### Solution

```go
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
```
