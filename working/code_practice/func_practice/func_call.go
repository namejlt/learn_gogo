package main

import "fmt"

func sumSub(param1, param2 int) (sum, sub int) {
	sum = param1 + param2
	sub = param1 - param2
	return sum, sub
}
func main() {
	p1 := 10
	p2 := 5
	summation, subtract := sumSub(p1, p2)
	fmt.Println(summation, subtract)
}
