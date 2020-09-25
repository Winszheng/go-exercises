package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//
func main() {
	//1.客户端基于tcp连接服务端
	conn, err := net.Dial("tcp", "localhost:1234")
	if err!= nil {
		log.Fatal("dial error: ", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	//2.客户端调用服务端提供的rpc服务
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(reply)
}
