[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/022-generate-parentheses/022-generate-parentheses.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next](https://github.com/albertopformoso/go-leetcode/blob/main/024-swap-nodes-in-pairs/024-swap-nodes-in-pairs.md)

# Merge K Sorted Lists

### Problem

https://leetcode.com/problems/merge-k-sorted-lists/

### Solution

```go
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}

	iteration, amount := 0, len(lists)-1
	for iteration < amount {
		midElement := (iteration + amount - 1) / 2
		for i := 0; i <= midElement; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[amount-i])
		}
		amount = (iteration + amount) / 2
	}

	return lists[0]
}

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
