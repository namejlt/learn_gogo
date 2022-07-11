package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "Good"
	greeting[1] = "Bad"
	greeting[2] = "Better"
	fmt.Println(greeting) //此时的切片为[Good Bad Better]

	greeting = append(greeting, "早上好") //此时的切片为[Good Bad Better 早上好 ]                  len=4,cap=5
	greeting = append(greeting, "中午好") //此时的切片为[Good Bad Better 早上好 中午好]             len=5,cap=5
	greeting = append(greeting, "下午好") //此时的切片为[Good Bad Better 早上好 中午好 下午好]       len=6,cap=10
	fmt.Println(len(greeting), cap(greeting))
	greeting = append(greeting, "晚上好") //此时的切片为[Good Bad Better 早上好 中午好 下午好 晚上好] len=7,cap=10
	fmt.Println(len(greeting), cap(greeting))

	fmt.Println(greeting[6])
	fmt.Println(len(greeting)) //切片具有自动扩容的机制：cap<1024时，成倍扩容。cap>1024时，每次增加1/4
	fmt.Println(cap(greeting))
}
