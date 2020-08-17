package main

import "fmt"

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		fmt.Print(num ,"^",nums[i],"=")
		num = num ^ nums[i]
		fmt.Println(num)
	}
	return num
}

//异或：相同为0，不同为1
//0 和 num 异或 都等于 num

func main(){
	a :=[]int{2,1,4,1,4}
	singleNumber(a)
}
