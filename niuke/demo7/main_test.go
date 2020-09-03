package demo7

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T){
	a :=[]int{1, 2, 3, 4, 5,6,7,8,9}
	root :=CreateTree(0,a)
	res :=threeOrders(root)
	fmt.Println(res)

}
