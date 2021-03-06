package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string) //声明一个map并初始化

	fmt.Println("here", words)

	sc := bufio.NewScanner(res.Body) //将http请求的响应体以指针的形式赋值给sc
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}

}
