package main

import "fmt"

//两个数组的交集
//输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
func intersect(nums1,nums2 []int) (res []int) {
	m1 :=make(map[int]int)
	for _,v:=range nums1{
		m1[v]++
	}
	for _,v :=range nums2{
		if m1[v] >0{
			m1[v]--
			res = append(res,v)
		}
	}
	return
}
func main() {
	a1 := []int{1,2,2,4,4,5}
	a2 :=[]int{4,3,1,2,2,2,0,5}
	fmt.Println(intersect(a1,a2))
}
