package strdemo3

import "strings"
/*
请你写一个函数StrToInt，实现把字符串转换成整数这个功能。
当然，不能使用atoi或者其他类似的库函数

你的函数应满足下列条件：
忽略所有行首空格，找到第一个非空格字符，可以是 ‘+/−’ 表示是正数或者负数，紧随其后找到最长的一串连续数字，将其解析成一个整数；
整数后可能有任意非数字字符，请将其忽略；
如果整数长度为0，则返回0；
如果整数大于INT_MAX(2^31 − 1)，请返回INT_MAX；如果整数小于INT_MIN(−2^31) ，请返回INT_MIN；
*/

func StrToInt(str string)int{
	max := 1<<31-1
	min := -1<<31
	if len(str) == 0{
		return 0
	}
	//去除首位空格
	str  = strings.Trim(str," ")
	k :=0
	//正负号标记
	sign :=0
	res :=0
	if str[k] == '+'{
		sign = 1
		k++
	}else if str[k] == '-'{
		sign = -1
		k++
	}
	for i:=k;i<len(str)&&(str[i]-'0'>=0 && str[i]-'0'<=9);i++{
		res = res*10 + int((str[i]-'0'))
	}
	if sign ==-1{
		res = -res
	}
	if res > max{
		return  res
	}
	if res < min{
		return min
	}
	return res
}

//编写去除首尾空格
func trim(str string)string{
	i,j :=0,len(str)-1
	for str[i] ==' '{
		i++
	}
	for str[j] ==' '{
		j--
	}
	return str[i:j+1]
}