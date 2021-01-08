package main

import (
	"fmt"
)

//用来判断一个字符串中汉字的个数
func hanzi(str string) int {
	bnewstr := []byte(str)
	rnewstr := []rune(str)
	return (len(bnewstr) - len(rnewstr)) / 2
}

func main() {
	s1 := "aaaa\\bbb\\ccc"
	s2 := `a\b\c
	cde
	fgh`
	fmt.Printf("%s\n", s1)
	fmt.Println(s2)
	newstr := "helloasda北京欢你爱几次"
	fmt.Printf("里面有%d汉字\n", hanzi(newstr))
}
