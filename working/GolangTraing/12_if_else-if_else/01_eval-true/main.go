package main

import "fmt"

func main() {

	if true { //if条件语句，如果为真，则执行下面代码块，否则跳过
		fmt.Println("This ran")
	}

	if false { //if条件语句，如果为假则跳过
		fmt.Println("This did not run")
	}
}
