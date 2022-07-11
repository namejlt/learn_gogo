package main

import "fmt"

func calculation(param1, param2 int) (sum, sub int) {
	sum = param1 + param2
	sub = param1 - param2
	return sum, sub
}
func main() {
	p1 := 10
	p2 := 5

	var f1 func(param1, param2 int) (sum, sub int)
	f1 = calculation

	/*f1, _ :=calculation(p1,p2)*/
	summation, subtract := f1(p1, p2)

	fmt.Println(summation, subtract)
}
