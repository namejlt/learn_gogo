package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	message := "zhang san is a name of dog."
	readMessage := strings.NewReader(message) //NewReader读取一个字符串并返回一个指针类型给readMessage
	//fmt.Println(readMessage)//*{zhang san is a name of dog. 0 -1}

	io.Copy(os.Stdout, readMessage) //os.Stdout和readMessage均是指针类型---Copy实现读写两个方法，并返回一个int64的written和error

	readMessage2 := bytes.NewBuffer([]byte(message)) //[]byte(message)将message转换成字节类型传给NewBuffer(buf []byte)
	//fmt.Println(readMessage2)

	io.Copy(os.Stdout, readMessage2)

	res, _ := http.Get("http://www.baidu.com") //Get请求返回http响应内容
	defer res.Body.Close()                     //关闭请求
	//fmt.Println(res)
	io.Copy(os.Stdout, res.Body) //copy将读到的内容写出来

}
