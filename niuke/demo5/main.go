package main

import "fmt"

func helper(a []int, left, right int) int {
	l, r := left, right
	mid := (l + r) / 2
	temp := a[mid]
	for l <= r {
		for r >= mid && a[r] >= temp {
			r--
		}
		if r >= mid {
			a[r], a[mid] = a[mid], a[r]
			mid = r
		}
		for l <= mid && a[l] <= temp {
			l++
		}
		if l <= mid {
			a[l], a[mid] = a[mid], a[l]
			mid = l
		}
	}

	return mid
}

func QuickSort(list []int, low, high int) {
	if high > low {
		//位置划分
		pivot := helperv1(list, low, high)
		//左边部分排序
		QuickSort(list, low, pivot-1)
		//右边排序
		QuickSort(list, pivot+1, high)
	}
}
func helperv1(data []int, left, right int) int {
	temp := data[left] //基准值
	mid := left

	for i := left + 1; i <= right; i++ {
		if temp > data[i] {
			data[i], data[mid+1] = data[mid+1], data[i]
			mid++
		}
	}
	data[left], data[mid] = data[mid], data[left]
	return mid

}

func main() {
	a := []int{2, 3, 1, 5, 3, 5, 8, 9, 0, 10, 4, 2}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
