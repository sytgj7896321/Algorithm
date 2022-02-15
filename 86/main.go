package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := new(ListNode)
	head2 := head
	tail := head
	for i := 0; i < 100; i++ {
		tail.Val = rand.Intn(12)
		tail.Next = new(ListNode)
		tail = tail.Next
	}
	tail.Next = nil

	for head2 != nil {
		fmt.Printf("%d ", head2.Val)
		head2 = head2.Next
	}
	fmt.Println()

	newList := partition(head, 9)
	for newList != nil {
		fmt.Printf("%d ", newList.Val)
		newList = newList.Next
	}
	fmt.Println()
}

func partition(head *ListNode, x int) *ListNode {
	small := new(ListNode)
	smallHead := small
	large := new(ListNode)
	largeHead := large
	cur := head
	for cur != nil {
		if cur.Val < x {
			small.Next = cur
			small = small.Next
		} else {
			large.Next = cur
			large = large.Next
		}
		cur = cur.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
