package main

import (
	"fmt"
)

func main() {
	student := []string{}    //声明一个切片并初始化,值为[]
	students := [][]string{} //声明一个数组类型切片并初始化,值为[]
	//students := [4][3]string{{"1", "2", "3"}, {"a", "b", "c"}, {}，{}} //
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	var stud []string //声明一个切片，未初始化，值为nil
	fmt.Println(stud == nil)
}
