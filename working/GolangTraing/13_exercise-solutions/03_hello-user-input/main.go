package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please enter your name: ") //Print 直接原样输出
	fmt.Scan(&name)                       //Scan需要在终端手动输出
	fmt.Println("Hello", name)            //打印输出
}
