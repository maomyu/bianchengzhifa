package main

import (
	"fmt"
	"strconv"
)

func helper(val int)int{
	max :=0
	i,j:=1,val-1
	for i <= j{
		a:=getSum2(i)
		b:=getSum2(j)
		if max < a + b{
			max = a+b
		}
		i++
		j--
	}
	return max
}
//各个数字之和
func getSum(val int)int{
	sum :=0
	str :=strconv.FormatInt(int64(val),10)
	r :=[]rune(str)
	for i:=0;i<len(r);i++{
		sum += int(r[i]-'0')
	}
	return sum
}

//各个数字之和 234
func getSum2(val int)int{
	sum :=0
	for val !=0{
		temp :=val % 10
		val = val /10
		sum +=temp
	}
	return sum
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}
func main(){
	T :=0
	fmt.Scan(&T)
	n :=0
	for i:=0;i<T;i++{
		fmt.Scan(&n)
		fmt.Println(helper(n))
	}
}