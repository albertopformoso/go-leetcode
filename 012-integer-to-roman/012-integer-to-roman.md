[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/011-integer-to-roman/011-integer-to-roman.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/013-roman-to-integer/013-roman-to-integer.md)

# Integer to Roman

### Problem

https://leetcode.com/problems/integer-to-roman/

### Solution

```go
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
```
