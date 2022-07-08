package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i) //无限循环，/死循环
		i++
	}
}
