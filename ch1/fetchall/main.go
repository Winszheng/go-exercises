package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// todo: 并发获取多个url并打印获取的字节数和分别花费的时间以及总时间
func main()  {
	start := time.Now()
	ch := make(chan string)	// 错误信息或正常执行信息
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Println("total: ", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch<-fmt.Sprint(err)
	}
	defer resp.Body.Close()
	n, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch<-fmt.Sprint(err)
		return
	}
	ch<-fmt.Sprint("fetch ",url," use ",time.Since(start).Seconds(),"s and get ",n," bytes")

	fmt.Println(flag.Arg(0))
	fmt.Println(flag.Args())
}

