package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)//（2的10次方）
	MB = 1 << (iota * 10)//（2的20次方）
	GB = 1 << (iota * 10)//（2的30次方）
	TB = 1 << (iota * 10)//（2的40次方）
)

func main() {
	fmt.Println("binary\t\tdecimal")//binary二进制，decimal十进制
	fmt.Printf("%b\t%d\t\n",KB,KB)
	fmt.Printf("%b\t%d\t\n",MB,MB)
	fmt.Printf("%b\t%d\t\n",GB,GB)
	fmt.Printf("%b\t%d\t\n",TB,TB)
}
