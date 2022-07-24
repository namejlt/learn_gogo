package main

import (
	"encoding/json"
	"os"
)

type man struct { //定义一个有四个字段的结构体man
	FirstName string
	LastName  string
	Age       int
	Number    int
}

func main() {
	p1 := man{"James", "Bond", 20, 7} //结构体实例化
	json.NewEncoder(os.Stdout).Encode(p1)
}
