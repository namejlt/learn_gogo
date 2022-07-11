package main

import "fmt"

func half(n int) (float64,bool) {
	return float64(n)/2,n%2==0//float64返回带有小数点的数

}

func main() {
	h,even:=half(5)
	fmt.Println(h,even)
}
