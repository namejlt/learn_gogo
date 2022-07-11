package main

import "fmt"

func main() {

	mySlice := []string{"星期一", "星期二", "星期三", "星期四", "星期五"} //string类型切片
	fmt.Println(mySlice)                                   //[星期一 星期二 星期三 星期四 星期五]

	otherSlice := []string{"星期六", "星期天"}

	mySlice = append(mySlice, otherSlice...) //给切片追加一个切片（这里和追加元素有区别）
	fmt.Println(mySlice)                     //[星期一 星期二 星期三 星期四 星期五 星期六 星期天]
}
