package main

import "fmt"

func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}                             //false  ||   false        ||  !false
                               //false       ||true
                                        //true
//&&也叫短路与：如果第一个条件为false，则第二个条件不会判断，最终结果为false
//||也叫短路或：如果第一个条件为true，则第二个条件不会判断，最终结果为true