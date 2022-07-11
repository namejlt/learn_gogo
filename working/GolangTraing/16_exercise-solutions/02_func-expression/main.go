package main

import "fmt"

func main() {

	half:= func(n int) (int,bool) {//函数可以作为入参
		return n/2,n%2==0
	}
	fmt.Println(half(5))

}
