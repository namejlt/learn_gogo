package main

import "fmt"

func main() {

	if true {
		fmt.Println("This ran") //布尔类型一般用于循环判断，当值为true时执行接下来的代码（块）
	}

	if false { //当值为false时执行接下来的代码（块）
		fmt.Println("This did not run")
	}
}
