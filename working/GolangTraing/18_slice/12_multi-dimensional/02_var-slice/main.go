package main

import "fmt"

func main() {

	var student []string //只有声明，没有初始化，值为nil
	var students []string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil, students == nil)
}
