package main

import (
	"fmt"
	"github.com/GoesToEleven/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName) //调用vis包下面的MyName变量，值为Todd。
	//fmt.Println(vis.yourName) //调用vis包下面的yourName变量，由于首字为小写，为不可导出变量，此句运行时会报错。
	vis.PrintVar() //调用vis包下面的PrintVar函数，该函数下面有两个变量MyName
}
