[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/013-roman-to-integer/013-roman-to-integer.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/015-3sum/015-3sum.md)

# Longest Common Prefix

### Problem

https://leetcode.com/problems/longest-common-prefix/

### Solution

```go
func longestCommonPrefix(strs []string) string {
    var longestPrefix string
    
    sort.Strings(strs)
    first := string(strs[0])
    last := string(strs[len(strs)-1])
    
    for i := 0; i < len(first); i++ {
        if string(last[i]) == string(first[i]) {
            longestPrefix += string(last[i])
        } else {
            break
        }
    }
    
    return longestPrefix
}
```
