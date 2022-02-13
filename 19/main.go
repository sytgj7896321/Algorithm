package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	test := new(ListNode)
	test.Val = 1
	test.Next = new(ListNode)
	test.Next.Val = 2
	test.Next.Next = new(ListNode)
	test.Next.Next.Val = 2
	test.Next.Next.Next = new(ListNode)
	test.Next.Next.Next.Val = 3
	test.Next.Next.Next.Next = new(ListNode)
	test.Next.Next.Next.Next.Val = 3
	test2 := removeNthFromEnd(test, 2)
	for test2 != nil {
		fmt.Println(test2.Val)
		test2 = test2.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	} else if head.Next == nil {
		return nil
	}
	l := head
	r := head
	for i := 0; i < n; i++ {
		if r == nil {
			break
		}
		r = r.Next
	}
	if r == nil {
		return head.Next
	}
	for {
		if r.Next == nil {
			l.Next = l.Next.Next
			break
		}
		r = r.Next
		l = l.Next
	}
	return head
}
