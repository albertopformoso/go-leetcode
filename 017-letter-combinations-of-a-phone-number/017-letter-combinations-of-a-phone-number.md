[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/016-3sum-closest/016-3sum-closest.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next]()

# Letter Combinations of a Phone Number

### Problem

https://leetcode.com/problems/letter-combinations-of-a-phone-number/

### Solution

```go
var letters = []string{0: "", 1: "", 2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz"}

func letterCombinations(digits string) (out []string) {

	if len(digits) > 0 {
		helper(&out, digits, "", 0)
	}

	return
}

func helper(out *[]string, digits, curr string, index int) {
	if index == len(digits) {
		*out = append(*out, curr)
		return
	}

	temp := letters[digits[index]-'0']
	for _, t := range temp {
		helper(out, digits, curr+string(t), index+1)
	}
}
```
