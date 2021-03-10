package main

import (
	"fmt"
)

func main() {
	// var array1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// //从数组定义切片切片的本质就是底层数组的封装
	// //切片从数组的0开始取到2但不包含2
	// sl1 := array1[0:2] //[1,2]
	// sl2 := array1[2:4] //[3,4]
	// sl3 := array1[5:7] //[6,7]
	// //容量是切片从开始到数组最右边的数量
	// fmt.Printf("长度:len(sl1)=%d 容量:cap(sl1)=%d\n", len(sl1), cap(sl1))
	// fmt.Printf("长度:len(sl2)=%d 容量:cap(sl2)=%d\n", len(sl2), cap(sl2))
	// fmt.Printf("长度:len(sl3)=%d 容量:cap(sl3)=%d\n", len(sl3), cap(sl3))
	// //必须用原来的切片接收append的返回值
	// sl3 = append(sl3, 1)
	// fmt.Printf("长度:len(sl3)=%d 容量:cap(sl3)=%d\n", len(sl3), cap(sl3))
	// //append如果原来容量放不下就会换一个新的底层数组
	// sl3 = append(sl3, 1)
	// fmt.Printf("长度:len(sl3)=%d 容量:cap(sl3)=%d\n", len(sl3), cap(sl3))
	// var s1 []int
	// //如果make的时候只设定长度的话 容量默认和长度相等s
	// s2 := make([]int, 4)
	// s_2 := make([]int, 0, 4)
	// s3 := []int{}
	// //s2和s3的定义方式都已经开辟了内存空间所以不为空
	// //因此通常判断切片是否为空的方式应该是len(s) == 0
	// fmt.Println(s1 == nil) //true
	// fmt.Println(s2 == nil) //false
	// fmt.Println(s3 == nil) //false
	// fmt.Println(len(s1), len(s2), len(s_2), len(s3))
	// ss1 := []int{1, 2, 3}
	// ss2 := ss1
	// ss3 := make([]int, 3)
	// //将ss1中的所有元素复制给ss3
	// copy(ss3, ss1)
	// ss1[0] = 0
	// fmt.Println(ss1, ss2, ss3)

	// ssss := []int{}
	// buff := new(bytes.Buffer)
	// if err := json.NewEncoder(buff).Encode(&ssss); err != nil {
	// 	fmt.Println(err)
	// }
	// var p []int
	// if err := json.NewDecoder(buff).Decode(&p); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(p)

	var kong []int = nil
	if len(kong) == 0 {
		fmt.Println("shi kong")
	}
	for i, v := range kong {
		fmt.Println(i, v)
	}

	var kong1 = []int{}
	if len(kong1) == 0 {
		fmt.Println("shi kong")
	}
	for i, v := range kong1 {
		fmt.Println(i, v)
	}

	fmt.Println(kong, kong1)
}
