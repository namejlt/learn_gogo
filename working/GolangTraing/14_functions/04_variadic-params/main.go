package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)//调用average函数并赋值给n
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)//打印sf类型
	var total float64
	for _, v := range sf {//循环求和sf元素之和
		total += v
	}
	return total / float64(len(sf))//求取数组sf内所有元素的平均值
}
