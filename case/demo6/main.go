package main

import (
	"fmt"
)
/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头
*/
func plusOne(digits []int)(res []int){
	jin :=1
	for i :=len(digits) - 1; i>=0;i--{
		res_temp :=digits[i] + jin
		if res_temp>=10{
			res = append(res,res_temp%10)
		}else{
			res = append(res,res_temp)
		}
		jin = res_temp /10
	}
	if jin >0{
		res = append(res,1)
	}
	res = resverNum(res)
	return
}
func resverNum(digits []int)[]int{
	i,j :=0,len(digits)-1
	for i < j{
		digits[i],digits[j] = digits[j],digits[i]
		i++
		j--
	}
	return digits
}

func plusOne1(digits []int) []int {
	i := 0
	for i = len(digits) - 1; i > 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}

	if i == 0 && digits[0] == 9 {
		digits[0] = 0
		// 说明上述循环没有提前终止
		digits = append(digits, 1)
		for j := len(digits) - 1; j > 0; j-- {
			digits[j], digits[j-1] = digits[j-1], digits[j]
		}
	} else if i == 0 {
		digits[0]++
	}
	return digits
}
func main(){
	a :=[]int{9,9,11,9}
	res :=plusOne(a)
	fmt.Println(res)
}
