package main

import "fmt"

func main() {

	if true && false { // 逻辑运算符&&也叫短路与：如果第一个条件为false，则第二个条件不会判断，最终结果为false（有假便为假）
		fmt.Println("This did not run")
	}

}
