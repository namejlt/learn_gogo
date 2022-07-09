package main

import "fmt"

func main() {
	var x [256]int //声明一个含有256个元素的数组

	fmt.Println(len(x))        //打印数组的长度
	fmt.Println(x[42])         //打印数组第43个元素的值
	for i := 0; i < 256; i++ { //通过循环给数组每个元素赋值
		x[i] = i
	} //[0,1,2,...,255]

	for i, v := range x {
		fmt.Printf("%v-%T-%b\n", v, v, v) //这里先打印，下面再判断
		if i > 50 {
			break //退出循环
		}
	}
}