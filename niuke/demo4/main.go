package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

//创建链表
func Create(data []int,index int)(head *ListNode){
	if index == len(data){
		return
	}
	head = &ListNode{data[index],nil}
	head.Next = Create(data,index+1)
	return
}

func mergeTwoLists( l1 *ListNode ,  l2 *ListNode ) *ListNode {
	head :=new(ListNode)
	p :=head
	for l1 !=nil || l2 !=nil{
		if l1 == nil{
			p.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil{
			p.Next = l1
			l1 = nil
			continue
		}
		if l1.Val >= l2.Val {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}else{
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		}
	}
	return head.Next
}

func Read(head *ListNode){
	p :=head
	for p !=nil{
		fmt.Print(p.Val," ")
		p = p.Next
	}
	fmt.Println()
}

func main(){
	a :=[]int{1,2,3,4,5,6,7,8,9}
	b :=[]int{2,3,5,11,12,13}
	l1 :=Create(a,0)
	l2 :=Create(b,0)
	Read(l1)
	Read(l2)
	//3 2 1
	h :=mergeTwoLists(l1,l2)
	Read(h)
}