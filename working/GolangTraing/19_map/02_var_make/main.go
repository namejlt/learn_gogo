package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string) //map必须初始化分配内存后才能被使用，用make来分配内存
	myGreeting["Tim"] = "Good morning."      //向map中添加数据
	myGreeting["Jenny"] = "Bonjour."         //map中的数据是以key-value的形式存在

	fmt.Println(myGreeting)
}
