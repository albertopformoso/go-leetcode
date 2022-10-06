package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	res := mergeTwoLists(list1, list2)
	fmt.Println(res)
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
