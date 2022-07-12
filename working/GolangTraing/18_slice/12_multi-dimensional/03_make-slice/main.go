package main

import "fmt"

func main() {
	student := make([]string, 35) //make给切片分配内存，相当于初始化
	students := make([][]string, 35)

	fmt.Println(len(students), cap(students)) //35 35

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil, students == nil) //初始化后值为空切片[]，而不是nil
}
