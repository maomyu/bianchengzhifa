package demo8

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T){
	a :=[][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(spiralOrder(a))
}
