package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type : %v name : %v kind : %v\n", v, v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("value : %v\n", v)
}

type person struct {
	name string
	age  int
}

func main() {
	// 获取值的类型
	var a float32 = 3.14
	reflectType(a)
	b := "nihao"
	reflectType(b)

	// 获取其他值的类型与底层类型
	var c *float32 = &a
	var d myInt = 2
	var e rune
	reflectType(c)
	reflectType(d)
	reflectType(e)
	f := person{
		name: "aaa",
		age:  18,
	}
	g := &person{
		name: "aaa",
		age:  18,
	}
	reflectType(f)
	reflectType(g)

	// 获取
}
