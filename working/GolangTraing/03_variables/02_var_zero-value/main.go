package main

import "fmt"

func main() {

	var a int
	var b string
	var c float64
	var d bool
	//输出go语言不同类型的默认值
	fmt.Printf("%v \n", a) //整形的默认值是0
	fmt.Printf("%v \n", b) //字符串类型的默认值是" "
	fmt.Printf("%v \n", c) //浮点型变量的默认值是0
	fmt.Printf("%v \n", d) //布尔值类型的默认值是false
	fmt.Println()
}
