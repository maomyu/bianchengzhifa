package main

import (
	"fmt"
	"github.com/yuwe1/bianchengzhifa/string/strdemo1"
	"github.com/yuwe1/bianchengzhifa/string/strdemo2"
	"github.com/yuwe1/bianchengzhifa/string/strdemo3"
)

func testlink(){
	for i := 1; i < 7; i++ {
		strdemo1.Add(strdemo1.GetHead(), i)
	}
	strdemo1.TravelLink(strdemo1.GetHead())
	strdemo1.LinkLeftShiftOne1(strdemo1.GetHead(),2)
	strdemo1.TravelLink(strdemo1.GetHead())
}
func teststrdemo2(){
	strdemo2.StringContain("abcdef","adcdr")
}
func teststrdemo3(){
	fmt.Println(strdemo3.StrToInt("123"))
	fmt.Println(strdemo3.StrToInt("   123434  1   "))
	fmt.Println(strdemo3.StrToInt("+13244545A"))
	fmt.Println(strdemo3.StrToInt("-23254 e"))
	fmt.Println(strdemo3.StrToInt("-232542353443365"))
}
func main() {

}
