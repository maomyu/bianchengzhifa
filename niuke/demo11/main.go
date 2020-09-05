package main

import "fmt"

/**
 *
 * @param x int整型
 * @return int整型
 */
func sqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid > x {
			r = mid - 1
		} else if mid*mid < x {
			l = mid + 1
		} else {
			return mid
		}
	}
	return r
}

func main() {
	a := 3
	fmt.Println(sqrt(a))
}
