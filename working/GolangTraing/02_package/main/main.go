package main

import (
	"fmt"
	"github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
	"github.com/GoesToEleven/GolangTraining/02_package/stringutil"
	//someAlias "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH")) //调用包stringutil内的Reverse函数，并传入参数为s="!oG ,olleH"
	fmt.Println(stringutil.MyName)
	fmt.Println(winniepooh.BearName)
}
