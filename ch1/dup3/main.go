package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// todo: 打印所有文件中每行的出现次数
func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	for _, filename := range filenames {
		text, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Panicln(err)
		}

		for _, v := range strings.Split(string(text),"\n") {
			counts[v]++
		}
	}

	for k, v := range counts {
		fmt.Println(k, v)
	}
}
