package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	fmt.Println(res)
	fmt.Println(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body) //ReadAll函数返回两个参数，一个是[]byte类型, 一个是error类型
	fmt.Println(bs)
	str := string(bs) //将字节类型转换成字符串类型
	fmt.Println(str)

}
