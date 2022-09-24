[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/001-two-sum/001-two-sum.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/003-longest-substring-without-repeating-characters/003-longest-substring-without-repeating-characters.md)

# Add Two Numbers

### Problem

https://leetcode.com/problems/add-two-numbers/

### Solution
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var carry int

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, val, v1, v2 int
	if l1 != nil { v1 = l1.Val }
	if l2 != nil { v2 = l2.Val }

	sum += v1 + v2 + carry
	carry = sum / 10
	val = sum % 10

	if l1 != nil { l1 = l1.Next } else { l1 = nil}
	if l2 != nil { l2 = l2.Next } else { l2 = nil}

	if l1 != nil || l2 != nil || carry != 0 {
		return &ListNode{val, addTwoNumbers(l1, l2)}
	}

	return &ListNode{val, nil}
}

```