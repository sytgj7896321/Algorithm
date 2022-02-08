package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := new(ListNode)
	b := new(ListNode)
	fmt.Println(mergeTwoLists(a, b))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	header := new(ListNode)
	newList := header
	for {
		if list1 != nil || list2 != nil {
			newList.Next = new(ListNode)
			newList = newList.Next
		} else {
			break
		}
		switch {
		case list1 == nil:
			newList.Val = list2.Val
			list2 = list2.Next
		case list2 == nil:
			newList.Val = list1.Val
			list1 = list1.Next
		case list1.Val <= list2.Val:
			newList.Val = list1.Val
			list1 = list1.Next
		case list2.Val < list1.Val:
			newList.Val = list2.Val
			list2 = list2.Next
		}
	}
	return header.Next
}
