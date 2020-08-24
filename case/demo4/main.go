package main

import "fmt"

func a(data []int,left,right int){
	if left < right{
		i,j :=left,right
		mid :=i
		for i <= j{
			for mid <=j && data[mid] <= data[j]{
				j--
			}
			if mid <=j{
				data[mid],data[j] = data[j],data[mid]
				mid = j
			}

			for mid >=i && data[mid] >= data[i]{
				i++
			}
			if mid >=i{
				data[mid],data[i] = data[i],data[mid]
				mid = i
			}

			a(data,left,mid-1)
			a(data,mid+1,right)
		}
	}
}

func main() {
	data :=[]int{-1,3,-4,5,1,3,5,23,5,7,3,2,10}
	a(data,0,len(data)-1)
	fmt.Println(data)
}
