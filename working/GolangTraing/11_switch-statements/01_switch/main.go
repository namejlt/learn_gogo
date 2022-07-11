package main

import "fmt"

func main() {

	switch "张三" { //switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
	case "李四":
		fmt.Println("得分80")
	case "王五":
		fmt.Println("得分90")
	case "张三":
		fmt.Println("得分95")
	default:
		fmt.Println("无法匹配")

	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
