package main

import "fmt"

func main() {
	// ln := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	// ln := &ListNode{1, nil}
	ln := &ListNode{1, &ListNode{2, nil}}
	res := removeNthFromEnd(ln, 2)
	fmt.Println(res)
}

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
