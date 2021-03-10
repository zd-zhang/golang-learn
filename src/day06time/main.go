package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	nowUnix := now.Unix()
	fmt.Println(nowUnix)
	newtime := now.Add(24 * time.Hour)
	fmt.Println(now, newtime)
	//格式化时间 把语言中时间对象转化为字符串类型时间
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("2006/01/02 03:04:05PM"))
	fmt.Println(now.Format("2006/01/02 03:04:05.000"))
	t, err := time.Parse(time.RFC3339, "2021-03-02T08:09:45Z")
	if err != nil {
		fmt.Printf("parse time err:%v", err)
	}
	fmt.Println(t, t.Unix())
	fmt.Println(t.UTC().Sub(now))

	time.Sleep(3 * time.Second)
	fmt.Println("sleep 3 seconds ")
}

func shiqu() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed err:%v\n", err)
		return
	}
	//按照东八区的时区和格式去解析一个字符串格式的时间
	t, err := time.ParseInLocation(time.RFC3339, "2021-02-03T08:12:34Z", loc)
	if err != nil {
		fmt.Printf("loacation to string err:%v", err)
		return
	}
	fmt.Println(t)
}

func main() {
	f1()
	shiqu()
}
