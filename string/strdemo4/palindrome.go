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