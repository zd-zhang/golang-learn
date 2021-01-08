package main

import "fmt"

//iota 1:遇到const就被重置为0 2:每有一行常量定义就加一
const (
	// 定义常量时如果不写后面的值那么默认与上一行相同
	a = iota + 1
	b //默认为 b = iota + 1
	c = 320
	d        // d等于320因为与上一行相等
	e = iota //e = 4
	f
)

var (
	s1 string = "s1"
	s2 string //不赋值那么s2为默认值""
)

const ()

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
}
