package main

import "fmt"

func main() {

	myGreeting := map[string]string{}   //声明一个map
	myGreeting["Tim"] = "Good morning." //向map中增加键值对
	myGreeting["Jenny"] = "Bonjour."    //向map中增加键值对

	fmt.Println(myGreeting)
}
