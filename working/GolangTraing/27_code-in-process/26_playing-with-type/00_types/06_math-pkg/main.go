package main

import (
	"fmt"
	"math"
)

func main() {
	up := 34.705945
	down := 34.405945
	fmt.Println(math.Floor(up + 0.5)) //Floor()函数  只舍不入 向下取整
	fmt.Println(math.Floor(down + 0.5))
}
