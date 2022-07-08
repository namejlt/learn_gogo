package main

import "fmt"

func getMaxNumber(number ...int) int {// ...+数据类型，表示函数可以传递任意数量的参数
	var largest int
	for _, value := range number {//for k,v:=range 参数{}，，，golang for循环常用方式
		if value>largest {
			largest=value//每次读到一个比largest大的value，就将此value赋值给largest,始终保持largest为最大的一个数
		}
	}
	return largest
}

func main() {
	maxNumber:=getMaxNumber(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println(maxNumber)
}
/*
FYI
For your code to also work with only negative numbers such as

greatest := max(-200 -700)

include this as your range statement
for i, v := range numbers {
	if v > largest || i == 0 {
		largest = v
	}
}

What does that code do?

The first time through the range loop
the index, i, will be zero
so largest will be set to the first number

Originally largest is set to the zero value for an int, which is zero

Zero would be greater than any negative number

if you only have negative numbers
you need largest to be something less than zero

Thanks to Ricardo G for this code improvement!
*/
