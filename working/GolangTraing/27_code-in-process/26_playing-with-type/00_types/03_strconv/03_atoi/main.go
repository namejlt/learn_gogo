package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("32") //将字符串类型的入参输出为整形类型的参数
	fmt.Println(i + 10)
}
