package main

import (
	"fmt"
)

func main(){
	n:=0
	fmt.Scan(&n)
	data :=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&data[i])
	}
	a := func(data []int)int{
		if len(data) == 0{
			return 0
		}

		dp :=make([]int,len(data))
		dp[0] = data[0]
		max :=data[0]
		for i:=1;i<len(data);i++{
			if dp[i-1]>=0{
				dp[i] = dp[i-1]+data[i]
			}else if dp[i-1]<0{
				dp[i] = data[i]
			}
			if max < dp[i]{
				max = dp[i]
			}
		}
		return max
	}
	fmt.Println(a(data))
}

/**动态规划
2 -1 3
dp[i-1] >=0 dp[i] = dp[i-1]+data[i]
dp[i-1] < 0 dp[i] = data[i]
 */