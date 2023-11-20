package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	for l != nil {
		if l.Next == nil {
			s += fmt.Sprintf("(%d)", l.Val)
		} else {
			s += fmt.Sprintf("(%d)→", l.Val)
		}
		l = l.Next
	}
	return s
}

func NewList(values ...int) *ListNode {
	var head *ListNode
	var current *ListNode

	for _, v := range values {
		node := &ListNode{Val: v}

		if head == nil {
			head = node
		} else {
			current.Next = node
		}
		current = node
	}

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var current *ListNode
	var carry int

	for l1 != nil || l2 != nil {
		var v1, v2 int

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		node := &ListNode{Val: v1 + v2 + carry}
		carry = 0

		if node.Val >= 10 {
			node.Val -= 10
			carry = 1
		}

		if result == nil {
			result = node
		} else {
			current.Next = node
		}

		current = node
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return result
}

func main() {
	a := NewList(2, 4, 3)
	b := NewList(5, 6, 4)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(addTwoNumbers(a, b))
}
