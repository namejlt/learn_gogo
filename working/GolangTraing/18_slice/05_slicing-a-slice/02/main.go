package main

import "fmt"

func main() {

	var results []int //定义一个整形切片
	fmt.Println(results)
	//0    1    2    3    4    5
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)       //[a b c g m z]
	fmt.Println(mySlice[2:4])  //取切片中元素下标[2,4)对应的值
	fmt.Println(mySlice[2])    //取切片中元素下标2对应的值
	fmt.Println("myString"[2]) //.........

}
