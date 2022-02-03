package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := new(ListNode)
	a.Val = 0
	p1 := a
	b := new(ListNode)
	b.Val = 0
	p2 := b
	for i := 0; i <= 200; i++ {
		p1.Val = i % 10
		p1.Next = new(ListNode)
		p1 = p1.Next
	}
	for i := 201; i < 400; i++ {
		p2.Val = i % 10
		p2.Next = new(ListNode)
		p2 = p2.Next
	}
	p := addTwoNumbers(a, b)
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	tail := head
	sum := 0
	for l1 != nil || l2 != nil || sum > 0 {
		tail.Next = new(ListNode)
		tail = tail.Next
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tail.Val = sum % 10
		sum /= 10
	}
	return head.Next
}
