package main

import "fmt"

func main() {

	myMap := map[int]string{
		0: "Golang",
		1: "Python",
		2: "Java",
		3: "PhP",
		4: "JS",
		5: "Ruby",
		6: "C++",
	}

	for key, val := range myMap { //遍历map并把键值对一一输出
		fmt.Println(key, " - ", val)
	}
	/*notice：map输出的键值对是无序的,每次输出的先后顺序都有可能不一样
		4  -  JS
		5  -  Ruby
		6  -  C++
		0  -  Golang
		1  -  Python
		2  -  Java
		3  -  PhP
	------------------
		1  -  Python
		2  -  Java
		3  -  PhP
		4  -  JS
		5  -  Ruby
		6  -  C++
		0  -  Golang
	*/

}
