package main

import "fmt"

func main() {
	intro := "Four 世"
	fmt.Printf("%T\n", intro) //打印类型
	fmt.Println(intro)        //打印值
	bs := []byte(intro)       //转换成字节
	fmt.Println(bs)           //打印字节
	fmt.Printf("%T\n", bs)    //bs的类型为[]uint8
	fmt.Println("*********")
	fmt.Printf("%d\n", bs) //[70 111 117 114 32 228 184 150]

	for _, v := range bs {
		fmt.Printf("%d\t\t %#x\t %b\n", v, v, v) //循环遍历bs 分别打印字节中元素的十进制 十六进制 二进制
	}
	fmt.Println("*********")
	y := 9999999999999999

	fmt.Printf("%d\t\t %#x\t %b\n", y, y, y) //打印字y中元素的  十进制 十六进制 二进制
	fmt.Println(&y)                          //取地址
	fmt.Sprint(y)                            //
	fmt.Println("*********")

	z := 'h'
	fmt.Printf("%T\n", z)
}
