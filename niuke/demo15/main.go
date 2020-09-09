package main

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

func mergeKLists( lists []*ListNode ) *ListNode {
	// write code here
	length := len(lists)

	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	mid := (length - 1) / 2
	leftList := mergeKLists(lists[:mid+1])
	rightList := mergeKLists(lists[mid+1 :])
	return mergeTwoLists(leftList, rightList)
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