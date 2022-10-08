package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	res := swapPairs(list1)
	fmt.Println(res)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node1 := &ListNode{head.Val, head.Next}
	node2 := &ListNode{head.Next.Val, head.Next.Next}

	node1.Next, node2.Next = swapPairs(node2.Next), node1
	
	return node2
}
