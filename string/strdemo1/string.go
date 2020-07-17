package strdemo1

import "fmt"

/*
给定一个字符串，要求把字符串前面的若干个字符移动到字符串的尾部，如把字符串“abcdef”前面的2个字符’a’和’b’移动到字符串的尾部，
使得原字符串变成字符串“cdefab”。请写一个函数完成此功能，要求对长度为n的字符串操作的时间复杂度为 O(n)，空间复杂度为 O(1)
*/

func LeftShiftOne1(str string, n int) string {
	if n >= len(str) || n <= 0 {
		return str
	}
	str = str[n:] + str[:n]
	return str
}

func LeftShiftOne2(str string, n int) string {
	for n > 0 && n < len(str) {
		str = str[1:] + str[:1]
		n--
	}
	return str
}

/*
三步反转
12    3456
21    6543
345612
*/
func resverString(str string) string {
	sr := []rune(str)
	i, j := 0, len(sr)-1
	for i <= j {
		sr[i], sr[j] = sr[j], sr[i]
		i++
		j--
	}
	return string(sr)
}
func LeftShiftOne3(str string, n int) string {
	if n >= len(str) || n <= 0 {
		return str
	}
	str1 := str[n:]
	str2 := str[:n]
	str1 = resverString(str1)
	str2 = resverString(str2)
	return resverString(str2 + str1)

}
func testString() {
	str := []string{"123456", "", "1", "12"}
	for _, v := range str {
		fmt.Println(LeftShiftOne1(v, 1))
	}
	fmt.Println("**********")
	for _, v := range str {
		fmt.Println(LeftShiftOne2(v, 1))
	}
	fmt.Println("**********")
	for _, v := range str {
		fmt.Println(LeftShiftOne3(v, 1))
	}
}

func LinkLeftShiftOne1(head *LinkNode, k int) {
	var pre *LinkNode = nil
	p :=head.next
	cur := head.next
	for i:=0;i<k;i++{
		temp :=cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}
	head.next = pre
	var pre1 *LinkNode = nil
	for cur!=nil{
		temp :=cur.next
		cur.next = pre1
		pre1 = cur
		cur = temp
	}
	p.next = pre1
}
