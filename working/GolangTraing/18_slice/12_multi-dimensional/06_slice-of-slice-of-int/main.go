package main

import "fmt"

func main() {

	slices := make([][]int, 0, 3) //初始化一个二维切片
	for i := 0; i < 3; i++ {      //外层循环，
		slice := make([]int, 0, 4) //初始化一个容量为四的int型切片
		for j := 0; j < 4; j++ {   //内层循环
			slice = append(slice, j) //每循环一次往slice中添加一个元素，循环四次
		}
		slices = append(slices, slice) //外层每循环一次，往slices中添加一个slice
	}
	fmt.Println(slices)
}
