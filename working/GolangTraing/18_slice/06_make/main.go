package main

import "fmt"

func main() {
	studentNumber := make([]int, 3) //切片长度和容量是都3
	studentNumber[0] = 7
	studentNumber[1] = 10
	studentNumber[2] = 15
	fmt.Println(studentNumber[0])
	fmt.Println(studentNumber[1])
	fmt.Println(studentNumber[2])

	greeting := make([]string, 3, 5) //切片长度为3，容量为5
	greeting[0] = "见到你"
	greeting[1] = "我"
	greeting[2] = "很高兴"

	fmt.Println(greeting[0], greeting[1], greeting[2])
}
