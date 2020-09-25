package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//首先是通过rpc.Dial拨号RPC服务，然后通过client.Call调用具体的RPC方法。
//在调用client.Call时，第一个参数是用点号链接的RPC服务名字和方法名字，
//第二和第三个参数分别我们定义RPC方法的两个参数。
func main() {
	//1.客户端基于tcp连接服务端
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err!= nil {
		log.Fatal("dial error: ", err)
	}

	var reply string
	//2.客户端调用服务端提供的rpc服务
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(reply)
}
