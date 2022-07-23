package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 5
	str := "Hello world " + strconv.Itoa(x) //strconv包下的Itoa()函数可以将整型数转换成字符串类型并返回字符串
	fmt.Println(str)
}
