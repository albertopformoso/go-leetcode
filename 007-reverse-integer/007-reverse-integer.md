[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/004-median-of-two-sorted-arrays/004-median-of-two-sorted-arrays.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/008-string-to-integer-(atoi)/008-string-to-integer-(atoi).md)

# Reverse Integer

### Problem

https://leetcode.com/problems/reverse-integer/

### Solution
```go
func reverse(num int) (revNum int) {
	for num != 0 {
		pop := num % 10
		num /= 10

        if revNum > math.MaxInt32/10 || (revNum == math.MaxInt32/10 && pop > 7) { return 0 }
		if revNum < math.MinInt32/10 || (revNum == math.MinInt32/10 && pop < -8) { return 0 }

		revNum = revNum*10 + pop
	}
	return
}
```