package demo7

import (
	"fmt"
)

type TreeNode struct {
	   Val int
	   Left *TreeNode
	   Right *TreeNode
}

func CreateTree(index int,data []int)(root *TreeNode){
	if index >= len(data){
		return
	}
	root = &TreeNode{data[index],nil,nil}
	root.Left =CreateTree(2*index+1,data)
	root.Right =CreateTree(2*index+2,data)
	return
}

//非递归
func PreRead(root *TreeNode)(res []int){
	//root left right
	p :=root
	var s []*TreeNode

	for p!=nil || len(s) > 0{
		for p != nil{
			res = append(res,p.Val)
			s = append(s,p)
			p = p.Left
		}
		if len(s)>0 {
			p = s[len(s)-1]
			s = s[:len(s)-1]
			p = p.Right
		}
	}
	return res
}

func MidRead(root *TreeNode)(res []int){
	p :=root
	s :=[]*TreeNode{}
	for p !=nil || len(s)>0{
		for p !=nil{
			s = append(s,p)
			p = p.Left
		}
		if len(s)>0{
			p = s[len(s)-1]
			res = append(res,p.Val)
			s = s[:len(s)-1]
			p = p.Right
		}
	}
	return res
}

func PostRead(root *TreeNode)(res []int){
	s := []*TreeNode{}
	pre :=root
	cur :=root
	if root == nil{
		return
	}
	s = append(s,root)
	for len(s) > 0{
		cur =s[len(s)-1]
		if (cur.Left == nil && cur.Right == nil) || (pre == cur.Left || pre == cur.Right){
			res = append(res,cur.Val)
			s = s[:len(s)-1]
			pre = cur
		}else{
			if cur.Right!=nil{s = append(s,cur.Right)}
			if cur.Left !=nil{s = append(s,cur.Left)}
		}
	}
	return res
}
func threeOrders( root *TreeNode ) [][]int {
	// write code here
	var res [][]int
	res = append(res,PreRead(root))
	fmt.Println(res)
	res = append(res,MidRead(root))
	fmt.Println(res)
	res = append(res,PostRead(root))
	fmt.Println(res)
    return  res
}


