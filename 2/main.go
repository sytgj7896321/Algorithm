package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i1, i2, sum1, sum2 := 0, 0, 0, 0
	p1, p2 := l1, l2
	for ; p1 != nil; i1++ {
		sum1 = sum1 + p1.Val*int(math.Pow(10, float64(i1)))
		p1 = p1.Next
	}
	for ; p2 != nil; i2++ {
		sum2 = sum2 + p2.Val*int(math.Pow(10, float64(i2)))
		p2 = p2.Next
	}
	sum := sum1 + sum2
	var iMax int
	if i1 > i2 {
		iMax = i1 - 1
	} else {
		iMax = i2 - 1
	}
	if sum >= int(math.Pow(10, float64(iMax+1))) {
		iMax++
	}
	var sumSliced []int
	for ; iMax >= 0; iMax-- {
		num := sum / int(math.Pow(10, float64(iMax)))
		sum = sum % int(math.Pow(10, float64(iMax)))
		sumSliced = append(sumSliced, num)
	}
	length := len(sumSliced)
	head := new(ListNode)
	head.Val = sumSliced[length-1]
	tail := head
	for i := length - 2; i >= 0; i-- {
		node := new(ListNode)
		node.Val = sumSliced[i]
		tail.Next = node
		tail = node
	}
	return head
}
