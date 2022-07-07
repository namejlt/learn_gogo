package main

import "fmt"

const metersToYards float64 = 1.09361//const用来定义常量——不会被改动的量

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")//Print 把输入对象默认为标准输出
	fmt.Scan(&meters)//手动在终端 输入 数值
	yards := meters * metersToYards//将米转换成码
	fmt.Println(meters, " meters is ",yards," yards.")
}
