package main

import (
	"demo-02/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//代码中最大的变化是用rpc.ServeCodec函数替代了rpc.ServeConn函数，传入的参数是针对服务端的json编解码器。
func main()  {
	//注册rpc服务到服务空间HelloService
	rpc.RegisterName("HelloService", new(service.HelloService))

	listen, err := net.Listen("tcp", ":1234")
	if err!= nil{
		log.Fatal("listen tcp error: ", err)
	}

	for{
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Accept err: ", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
//nc -l 1234 作为server端启动一个tcp的监听，监听的是1234端口
//netcat 用于创建网络连接