package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// todo: 检测每行出现的次数并输出
// 输入是标准输出或者文件
// 无论是标准输入还是文件，实际都被当成文件处理(*os.File)
// so实际要处理的只是输入的参数
func main()  {
	files := os.Args[1:]
	// stdin
	if len(files) == 0 {
		countline(os.Stdin)
	}else{
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil{
				log.Println(err)
			}
			countline(f)
		}
	}

}

func countline(file *os.File)  {
	input := bufio.NewScanner(file)	// 处理每行输入
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Println(k,v)
	}
}