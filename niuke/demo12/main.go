package main

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	p := new(ListNode)
	p.Next = head
	count := 0
	p1 := p.Next
	for p1 != nil {
		count++
		p1 = p1.Next
	}
	p2 := p
	for i := 0; i < count-n; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return p2.Next
}

func Read(head *ListNode) {
	p := head
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
}

//借助快慢指针
func removeNthFromEndv2(head *ListNode, n int) *ListNode {
	// write code here
	res := head
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// n == N
	if fast == nil {
		return res.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return res
}
func main() {
	data := []int{1}
	head := CreateLink(data, 0)
	head = removeNthFromEnd(head, 1)
	Read(head)
}
