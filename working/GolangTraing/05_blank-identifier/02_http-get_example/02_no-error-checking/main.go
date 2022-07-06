package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res,_:=http.Get("https://www.baidu.com/55555")//这里不考虑可能发生的错误
	page,_:=ioutil.ReadAll(res.Body)//用下划线_代替返回值err
	_ = res.Body.Close()
	//1
	fmt.Printf("%s",page)
	fmt.Printf("%s",page)
	fmt.Printf("%s",page)
	fmt.Printf("%s",page)
}
