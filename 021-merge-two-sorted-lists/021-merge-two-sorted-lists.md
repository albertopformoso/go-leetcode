[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/020-valid-parentheses/020-valid-parentheses.md) |) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/022-generate-parentheses/022-generate-parentheses.md)

# Merge Two Sorted Lists

### Problem

https://leetcode.com/problems/merge-two-sorted-lists/

### Solution

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, list2),
		}
	}

	return &ListNode{
		Val:  list2.Val,
		Next: mergeTwoLists(list1, list2.Next),
	}
}
```
