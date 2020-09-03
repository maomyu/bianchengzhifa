package demo6

import "fmt"

func findKthv1( a []int ,  n int ,  K int ) int {
	// 借助选择排序
	for i := 0; i < K; i++ {
		for j :=i+1;j< n;j++{
			if a[i] < a[j]{
				 a[i],a[j] = a[j],a[i]
			}
		}
		fmt.Println(a)
	}
	return a[K-1]
}

func finxkthv2(a []int ,  n int ,  K int)int{
	//二分查找

	return findhlper(a,0,n-1,K)
}

func findhlper(a []int,l,r int,k int)int{
	if l == r{
		return a[l]
	}
	mid :=helperv2(a,l,r)
	if mid == k-1{
		return a[mid]
	}else if mid > k-1{
		return findhlper(a,l,mid-1,k)
	}else{
		return findhlper(a,mid+1,r,k)
	}
}
func helperv2(a []int,l,r int)int{
	v :=a[l]
	mid :=l
	for i :=mid+1;i<=r;i++{
		if a[i] > v{
			a[mid+1],a[i] = a[i],a[mid+1]
			mid++
		}
	}
	a[l],a[mid] = a[mid],a[l]
	return mid
}

func quickSort(a []int,left,right int){
	if left <  right{
		mid  :=helperv2(a,left,right)
		quickSort(a,left,mid-1)
		quickSort(a,mid+1,right)
	}
}

