package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLink(data []int, index int) (head *ListNode) {
	if index >= len(data) {
		return
	}
	head = &ListNode{data[index], nil}
	head.Next = CreateLink(data, index+1)
	return
}
func Read(head *ListNode) {
	p := head
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
}

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	head1 = reverse(head1)
	head2 = reverse(head2)
	newh := new(ListNode)
	p := newh
	jin := 0
	for head1 != nil || head2 != nil || jin > 0 {
		if head1 != nil {
			jin += head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			jin += head2.Val
			head2 = head2.Next
		}
		p.Next = &ListNode{jin % 10, nil}
		p = p.Next
		jin = jin / 10
	}
	return reverse(newh.Next)
}

func reverse(h *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	for h != nil {
		tmp := h.Next
		h.Next = dummy.Next
		dummy.Next = h
		h = tmp
	}
	return dummy.Next
}

func main() {
	s1, s2 := []int{9, 9, 9, 7}, []int{6, 3}
	h1, h2 := CreateLink(s2, 0), CreateLink(s1, 0)
	h := addInList(h1, h2)
	Read(h)
}
