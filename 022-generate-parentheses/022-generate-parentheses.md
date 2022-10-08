[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/021-merge-two-sorted-lists/021-merge-two-sorted-lists.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/023-merge-k-sorted-lists/023-merge-k-sorted-lists.md)

# Generate Parentheses

### Problem

https://leetcode.com/problems/generate-parentheses/

### Solution

```go
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
```
