package main

import "fmt"

func main() {
	li := ListNode{Val: 5}
	obj := Constructor(&li)
	param_1 := obj.GetRandom()
	fmt.Println(param_1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	values []int
}

func Constructor(head *ListNode) Solution {
	var s Solution
	for head != nil {
		s.values = append(s.values, head.Val)
		head = head.Next
	}
	return s
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	rand.New(
	return this.values[0]
}
