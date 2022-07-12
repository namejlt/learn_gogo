package main

import (
	"fmt"
)

func main() {

	var records [][]string //定义一个二维切片
	// student 1
	s1 := make([]string, 4) //初始化一个容量为4的string类型切片s1
	s1[0] = "张三"            //给s1赋值
	s1[1] = "李四"
	s1[2] = "100.00"
	s1[3] = "74.00"

	records = append(records, s1) //将s1追加到records中

	s2 := make([]string, 4) //初始化一个容量为4的string类型切片S2
	s2[0] = "王五"
	s2[1] = "赵六"
	s2[2] = "92.00"
	s2[3] = "96.00"

	records = append(records, s2) //将s2追加到records中

	fmt.Println(records) //打印
}
