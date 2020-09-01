package lru

import (
	"fmt"
	"testing"
)

func TestLruv1(t *testing.T) {
	result :=[][]int{
		{1,1,1},{1,2,2},{1,3,2},{2,1},{1,4,4},{2,2},
	}
	fmt.Println(LRUV1(result,3))
}

