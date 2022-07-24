package main

import "fmt"

func main() {

	mySlice := []string{"星期一", "星期二"} //string类型切片
	fmt.Println(mySlice)              //[星期一 星期二]

	otherSlice := []string{"星期三", "星期四", "星期五"}

	mySlice = append(mySlice, otherSlice...) //给切片追加一个切片（这里和追加元素有区别）
	fmt.Println(mySlice)                     //[星期一 星期二 星期三 星期四 星期五]
	fmt.Println(mySlice[:2])                 //[星期一 星期二]
	fmt.Println(mySlice[3:])                 //[星期四 星期五]

	mySlice = append(mySlice[:2], mySlice[3:]...) //在原切片的基础上先进行截取mySlice[:2]，再给其append一个切片
	fmt.Println(mySlice)                          //[星期一 星期二 星期四 星期五]

}
