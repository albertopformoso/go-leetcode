[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/018-4sum/018-4sum.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/020-valid-parentheses/020-valid-parentheses.md)

# Remove Nth Node From End of List

### Problem

https://leetcode.com/problems/remove-nth-node-from-end-of-list/

### Solution

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := length(head)
	position := len - n
	if position == 0 {
		head = head.Next
	}

	return remove(head, 0, position)
}

func length(head *ListNode) int {
	if head == nil {
		return 0
	}
	return length(head.Next) + 1
}

func remove(head *ListNode, currPos, targetPos int) *ListNode {
	if head != nil {
		if currPos == targetPos-1 {
			head.Next = head.Next.Next
			return &ListNode{head.Val, remove(head.Next, currPos+1, targetPos)}
		}

		return &ListNode{head.Val, remove(head.Next, currPos+1, targetPos)}
	}

	return nil
}
```
