[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/012-integer-to-roman/012-integer-to-roman.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next]()

# Roman to Integer

### Problem

https://leetcode.com/problems/roman-to-integer/

### Solution

```go
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
```
