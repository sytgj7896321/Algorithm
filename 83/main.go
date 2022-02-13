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
	test.Next.Val = 1
	test.Next.Next = new(ListNode)
	test.Next.Next.Val = 2
	test.Next.Next.Next = new(ListNode)
	test.Next.Next.Next.Val = 3
	test.Next.Next.Next.Next = new(ListNode)
	test.Next.Next.Next.Next.Val = 3
	var test2 *ListNode
	deleteDuplicates(test)
	deleteDuplicates(test2)
	for test != nil {
		fmt.Println(test.Val)
		test = test.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := head
	r := head.Next
	for {
		if l.Val != r.Val {
			l.Next = r
			l = l.Next
			if r.Next != nil {
				r = r.Next
			}
		} else {
			if r.Next == nil {
				if l.Val == r.Val {
					l.Next = nil
					break
				} else {
					l.Next = r
					break
				}
			}
			r = r.Next
		}
	}
	return head
}
