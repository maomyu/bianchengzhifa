package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func C(head *Node,val int){
	if head.next == nil{
		head.next = &Node{val,nil}
		return
	}
	p :=head
	for p.next != nil{
		p = p.next
	}
	p.next = &Node{val,nil}
}

func Read(root *Node){
	if root == nil{
		return
	}
	p :=root.next
	for p!=nil{
		fmt.Print(p.val," ")
		p = p.next
	}
}

func helper(head *Node,k int){
	pre :=head
	cur :=pre.next
	for i:=0;i<k-1;i++{
		pre = pre.next
		cur = cur.next
	}
	pre.next = cur.next
}

func main(){
	head := new(Node)
	var a,k int
	fmt.Scan(&a,&k)
	if k > a{
		return
	}
	data :=make([]int,a)
	for i:=0;i<len(data);i++{
		fmt.Scan(&data[i])
	}
	for i:=0;i<len(data);i++{
		C(head,data[i])
	}
	helper(head,k)
	Read(head)
}
