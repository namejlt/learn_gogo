package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {//res表示来自HTTP请求的响应。 一旦收到响应头，客户端和传输将从服务器返回响应。读取正文字段时，响应正文按需流式传输。
	res, err := http.Get("https://www.baidu.com/") //http包下的Get方法向指定的url发起请求，并返回两个参数res和err
	if err != nil {
		log.Fatal(err) //Log包下的Fatal函数相当于print()，打印输出的作用
	}

	page, err := ioutil.ReadAll(res.Body)//ioutil包下的ReadAll函数读取httpGet请求的响应body（res.Body）并返回给[]byte类型的page
	_ = res.Body.Close()//当发起一个http请求后，需要手动关闭此请求。如果 response 不关闭，会导致内存泄漏。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)//将[]byte类型的page以string类型格式化输出
}
