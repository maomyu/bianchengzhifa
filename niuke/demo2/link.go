package main

import "fmt"

type ListNode struct{
    Val int
    Next *ListNode
}

//递归创建
func CreateLink( index int ,data []int)(head *ListNode){
	if len(data) <= index {
		return
	}
	head = &ListNode{data[index], nil}
	head.Next = CreateLink( index+1,data)
	return head
}

func CreateLinkV2(head *ListNode,data []int){
	var index int
	p :=head
	for index <len(data){
		p.Next = &ListNode{data[index],nil}
		p = p.Next
		index++
	}
}

//head->1->2->3->4->5->6->7->8->9
func ReverseList(pHead *ListNode ) *ListNode {
	var pre ,cur *ListNode
	cur  = pHead
	for cur !=nil{
		temp :=cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func Read(head *ListNode){
	p :=head
	for p !=nil{
		fmt.Print(p.Val," ")
		p = p.Next
	}
	fmt.Println()
}

func main() {
	data :=[]int{1,2,3,4,5,6,7,8,9}
	var head *ListNode = new(ListNode)
	CreateLinkV2(head, data)
	Read(head)
	head =CreateLink(0,data)
	Read(head)
	head = ReverseList(head)
	Read(head)
}