package strdemo4


func IspalindRome(str string)bool{
	i,j :=0,len(str)-1
	for i <=j{
		if str[i] != str[j]{
			return false
		}
	}
	return true
}
//abaaba
//aabaa
func IsPalindrome2(str string)bool{
	var right, left int
	l :=len(str)
	if l ==1{
		return true
	}
	if l & 1 ==0{
		left ,right = l/2-1,l/2
	}else{
		left, right =l/2-1,l/2+1
	}
	for left>=0 && right<l{
		if str[left]!=str[right]{
			return false
		}
		left--
		right++
	}
	return true
}