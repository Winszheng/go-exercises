package main

import (
	"bufio"
	"fmt"
	"os"
)
// input:
//aaa
//aaa
//aaa
//bbb
//ccc
//ddd

// output:
//aaa   3
//bbb   1
//ccc   1
//ddd   1

func main()  {
	counts := make(map[string]int)
	// 从标准输入中读取内容
	input := bufio.NewScanner(os.Stdin)
	// 每次执行就读取一行，并返回是否成功，直到没得读取
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Println(line, " ", n)
	}

}
