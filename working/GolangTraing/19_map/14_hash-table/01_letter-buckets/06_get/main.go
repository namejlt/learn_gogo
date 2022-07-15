package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	//Get方法向一个URL发出请求，返回该页面上的所有内容或者返回一个错误。
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close() //Get方法发出请求后，需要关闭此请求，防止内存泄露
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs) //格式化打印输出
	fmt.Println(res)     //这里返回的是http请求的结果，包括协议，请求内容类型，内容长度，日期，服务器等
	/*
		{200 OK 200 HTTP/1.1 1 1
		map[Accept-Ranges:[bytes]
		Content-Length:[1256167]
		Content-Type:[text/plain]
		Date:[Thu, 14 Jul 2022 11:03:52 GMT]
		Last-Modified:[Thu, 22 Jun 2000 04:00:00 GMT]
		Server:[Apache]]
		0xc000026180
		1256167
		[]
		false
		false
		map[]
		0xc00015c100
		0xc00013c2c0}
	*/

}
