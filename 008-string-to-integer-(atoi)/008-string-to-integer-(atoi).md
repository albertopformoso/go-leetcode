[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/007-reverse-integer/007-reverse-integer.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/009-palindrome-number/009-palindrome-number.md)

# String To Integer (atoi)

### Problem

https://leetcode.com/problems/string-to-integer-atoi/

### Solution
```go
func myAtoi(s string) (out int) {
	var i int
	var sign int = 1

	s = strings.TrimSpace(s)

	if len(s) == 0 { return out }
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for ; i<len(s); i++ {
		if sign*out >= math.MaxInt32 { return math.MaxInt32 }
		if sign*out <= math.MinInt32 { return math.MinInt32 }

		if s[i] < '0' || s[i] > '9' { break }

		out = out*10 + (int(s[i]) - '0')
	}

	if sign*out >= math.MaxInt32 { return math.MaxInt32 }
	if sign*out <= math.MinInt32 { return math.MinInt32 }

	return out * sign
}
```