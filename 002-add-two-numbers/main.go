package main

import (
	"fmt"
	"strings"
	"strconv"
)

type ListNode struct {
	Val 	int
	Next 	*ListNode
}

var carry int

func main() {
	l1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}

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

func (ln *ListNode) String() string {
	sb := strings.Builder{}
	ln.printNode(&sb, ln)
	return sb.String()
}

func (ln ListNode) printNode(sb *strings.Builder, root *ListNode) {
	if root == nil { return }

	sb.WriteString(strconv.Itoa(root.Val) + " ")
	ln.printNode(sb, root.Next)
}