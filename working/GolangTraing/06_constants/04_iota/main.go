package main

import "fmt"

const (
	_ = iota	 //0
	b = iota * 10//10
	c = iota * 10//20
)

func main() {
fmt.Println(b,c)
}
