package main

import "fmt"

func main() {
	var x [58]string//定义一个具有58个元素的string类型数组

	for i := 65; i <=122 ; i++ {//i从65自增到122
		x[i-65]=string(i)
		//i=65时，x[0]=string(65),string(65)为字符A
		//i=66时，x[1]=string(66),string(66)为字符B
		//...
		//i=122时，x[57]=string(122),string(122)为字符z
	}
	fmt.Println(x)
	fmt.Println(x[1])
	/*var x [58]string

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])*/
}
