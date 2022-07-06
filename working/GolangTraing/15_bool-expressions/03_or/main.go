package main

import "fmt"

func main() {

	if true || false { //逻辑运算符||也叫短路或：如果第一个条件为true，则第二个条件不会判断，最终结果为true（有真便为真）
		fmt.Println("This ran")
	}
}
