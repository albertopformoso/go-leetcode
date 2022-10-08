[Previous](https://github.com/albertopformoso/go-leetcode/blob/main/023-merge-k-sorted-lists/023-merge-k-sorted-lists.md) | [Home](https://github.com/albertopformoso/go-leetcode) | [Next]()

# Swap Nodes in Pairs

### Problem

https://leetcode.com/problems/swap-nodes-in-pairs/

### Solution

```go
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node1 := &ListNode{head.Val, head.Next}
	node2 := &ListNode{head.Next.Val, head.Next.Next}

	node1.Next, node2.Next = swapPairs(node2.Next), node1
	
	return node2
}
```
