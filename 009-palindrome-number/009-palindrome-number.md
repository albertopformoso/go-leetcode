[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/008-string-to-integer-(atoi)/008-string-to-integer-(atoi).md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/010-regular-expression-matching/010-regular-expression-matching.md)

# Palindrome Number

### Problem

https://leetcode.com/problems/palindrome-number/

### Solution

#### Not using strings
```go
func isPalindrome(x int) bool {
	var reversed int
	var original int = x

	if x < 0 { return false }

	for x > 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}

	return original == reversed
}
```

#### Using strings
```go
func isPalindrome(x int) bool {
	strNum := strconv.Itoa(x)
	revStrNum := reverse(strNum)

	fmt.Println(strNum, revStrNum)
	return strNum == revStrNum
}

func reverse(s string) string {
	var rev string
	for i := len(s)-1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}
```
