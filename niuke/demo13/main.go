package demo13

import "fmt"

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

func detectCycle(head *ListNode) *ListNode {
	// write code here
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			p := head
			for p != p1 {
				p = p.Next
				p1 = p1.Next
			}
			return p1
		}
	}
	return nil
}
