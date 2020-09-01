package lru

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
//[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]
/*
11
11 22
11 22 32
22 32 11

默认最新的为尾部，需要淘汰的为头部
*/
func LRUV1( operators [][]int ,  k int ) []int {
	// write code here
	keys := make([]int,0,k)
	values := make([]int,0,k)
	out := make([]int,0,len(operators))
	for _,o :=range operators {
		if o[0] == 1 {
			if len(keys) == k {
				keys = keys[1:]
				values = values[1:]
			}
			keys = append(keys, o[1])
			values = append(values, o[2])
		} else if o[0] == 2 {
			var i int
			for i=0;i<k;i++ {
				if i >= len(keys) {
					break
				}
				if o[1] == keys[i] {
					break
				}
			}
			if i==k || i>=len(keys) {
				out = append(out, -1)
			} else {
				out = append(out, values[i])
				if i+1 < k {
					keys = append(keys[0:i],append(keys[i+1:],keys[i])...)
					values = append(values[0:i],append(values[i+1:],values[i])...)
				}
			}
		}
	}
	return out
}

