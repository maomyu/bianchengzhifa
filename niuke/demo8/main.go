package demo8

//螺旋矩阵的输出
func spiralOrder( matrix [][]int ) []int {
	var res []int
	if len(matrix) == 0{
		return res
	}
	// write code here

	left,right :=0,len(matrix[0])-1
	top,bottom :=0,len(matrix)-1
	for left <= right && top <=bottom{
		//往右走
		for i:=left;i <= right;i++ {
			res = append(res,matrix[top][i])
		}
		top++
		//往下走
		for i :=top;i <= bottom;i++{
			res = append(res,matrix[i][right])
		}
		right--
		//往左走
		if top <= bottom{
			for i :=right; i >= left;i--{
				res = append(res,matrix[bottom][i])
			}
		}
		bottom--
		if left <= right{
			//往上走
			for i:=bottom ;i >= top;i--{
				res = append(res,matrix[i][left])
			}
		}

		left++
	}
	return res
}
/**
a :=[][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
a :=[][]int{
	{2,3},
}
 */

/**
从左到右 （left,right）-> (left++,right) top++代表当前第一层遍历完成
从上到下 （top,right）-> (top++,right) right-- 代表当前右边遍历完成
从右到左  (right,left) ->(right--,left) 同时要保证top和bottom满足条件
.......
 */