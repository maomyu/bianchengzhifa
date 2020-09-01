package lru


func LRUv2(operators [][]int ,  k int)[]int{
	keys :=make([]int,0,k)
	values :=make([]int,0,k)
	out :=make([]int,0,len(operators))
	for _,v :=range operators{
		if v[0] == 1{
			//set
			if len(keys) == k{
				keys = keys[1:]
				values=values[1:]
			}
			keys = append(keys,v[1])
			values = append(values,v[2])
		}else if v[0] == 2{
			sign :=0
			//get
			targetK := v[1]
			for i:=0;i<len(keys);i++{
				if keys[i] == targetK{
					out = append(out,values[i])
					//move
					if i+1<k{
						keys = append(keys[0:i],append(keys[i+1:],keys[i])...)
						values = append(values[0:i],append(values[i+1:],values[i])...)
					}
					sign = 1
					break
				}
			}
			if sign == 0 {
				out = append(out,-1)
			}
			sign = 0
		}
	}
	return out
}