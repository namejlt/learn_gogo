package main

import "fmt"

func main() {
	fmt.Println(greeting("Alex","Lucy"))
}
func greeting(name,describe string) (string,string) {//2个入参name，2个输出参数
	return fmt.Sprint(name," ",describe),fmt.Sprint(describe," ",name)//有两个输出参数，必须return两个string类型的值
}
