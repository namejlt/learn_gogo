package main

import "fmt"

func add(slice ...int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}
func main() {
	fmt.Println("1+2+3+4+5=", add(1, 2, 3, 4, 5))
}
