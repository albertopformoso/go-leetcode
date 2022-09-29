package main

import (
	"fmt"
	"sort"
)

func main() {
	res := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(res)
}

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
