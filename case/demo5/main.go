package main

import "fmt"

func Zichuan(str string)[][]string{
	res :=make([][]string,0)
	path :=make([]string,0)
	helper(str,0,len(str),path,&res)
	return res
}
func helper(str string,start int,l int,path []string,res *[][]string){
	if start == l{
		*res = append(*res,path)
		return
	}
	for i:=start;i<l;i++{
		path = append(path,str[start:i+1])
		helper(str,i+1,l,path,res)
		path = path[:len(path)-1]
	}
}
func main(){
	str :="123"
	res :=Zichuan(str)
	fmt.Println(res)
}