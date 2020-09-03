package demo6

import (
	"fmt"
	"testing"
)

func TestFindk(t *testing.T) {
	a :=[]int{1, 2, 3,4,5,6,7,8,9}
	fmt.Println(findKthv1(a,len(a),3))
}

func TestQuick(t *testing.T){
	a :=[]int{5, 9, 2,3,7,0,7,4,1}
	quickSort(a,0,len(a)-1)
	fmt.Println(a)
}

func TestFindk2(t *testing.T) {
	a :=[]int{1, 9, 2,3,7,0,6,4,5}
	fmt.Println(finxkthv2(a,len(a),3))
}