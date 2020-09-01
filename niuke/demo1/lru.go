package lru

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

var target [][]int

func LRU( operators [][]int ,  k int )(res []int) {
	var index int
	// write code here
	for i:=0;i<len(operators);i++ {
		cur :=operators[i]
		if cur[0] == 1{
			//set
			target = Set(target,cur)
			index++
			if index > k{
				target = ReMove(target)
				//移除
				index--
			}
		}else {
			r ,result:=Get(target,cur)
			target = result
			res = append(res,r)
		}
	}
	return res
}

func Set(target [][]int,cur []int)(res [][]int){
	sign :=0
	//judge exist
	for i:=0;i<len(target);i++{
		if cur[1] == target[i][0]{
			res = append(res,cur[1:])
			res = append(res,target[:i]...)
			res = append(res,target[i+1:]...)
			sign = 1
		}
	}
	if sign == 0{
		res = append(res, cur[1:])
		res = append(res,target[:]...)
	}
	return res
}

func ReMove(res [][]int)[][]int{
	return res[:len(res)-1]
}

func Get(target [][]int,cur []int)(r int,res [][]int){
	for i:=0;i<len(target);i++{
		if cur[1] == target[i][0]{
			temp := target[i][1]
			res = append(res,cur[1:])
			res = append(res,target[:i]...)
			res = append(res,target[i+1:]...)
			return temp,res
		}
	}
	return -1,res
}
