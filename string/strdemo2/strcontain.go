package strdemo2

import "fmt"

/*
给定两个分别由字母组成的字符串A和字符串B，字符串B的长度比字符串A短。
请问，如何最快地判断字符串B中所有字母是否都在字符串A里？
*/
func StringContain(str1,str2 string)bool{
	hash :=0
	for i:=0;i<len(str1);i++{
		fmt.Println(hash,"和",1<<(str1[i]-'A'),"做或运算")
		hash |=(1<<(str1[i]-'A'))
		fmt.Println("得到hash值：",hash)
	}
	fmt.Println()
	for i :=0;i<len(str2);i++{
		fmt.Println(hash ,"和", (1 << (str2[i]-'A')),"做与运算，得到hash值：",(hash & (1 << (str2[i]-'A'))))
		if ((hash & (1 << (str2[i]-'A')))==0){
			return false
		}
	}
	return true
}

/*
或运算：有一个为1结果就为1
	val |= 代表是对每一个字符进行合并，比如a为00001，b为00010
    ab =  a ｜ b =00011
与运算：每一位两者都为1，则结果为1，否则为0
	利用与运算判断是否有重复，比如，如果出现重复，则结果不为0，如果没有任何重复，其结果为0

*/