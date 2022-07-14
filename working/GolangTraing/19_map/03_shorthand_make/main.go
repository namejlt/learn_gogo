package main

import "fmt"

func main() {

	myGreeting := make(map[string]string) //简短声明和变量声明一样
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
