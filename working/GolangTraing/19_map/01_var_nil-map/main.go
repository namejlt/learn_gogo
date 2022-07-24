package main

import "fmt"

func main() {

	var myGreeting map[string]string //声明一个map

	fmt.Println(myGreeting) //map的默认值为map[]

	fmt.Println(myGreeting == nil) //map[]其实就是nil
}
