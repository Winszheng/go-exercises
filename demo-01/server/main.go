package main

import (
	"demo-01/service"
	"log"
	"net"
	"net/rpc"
)

func main()  {
	//1.注册grpc服务
	//其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	//所有注册的方法会放在“HelloService”服务空间之下。然后我们建立一个唯一的TCP链接，
	//并且通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	rpc.RegisterName("HelloService", new(service.HelloService))

	//2.创建网络监听器
	//典型做法：net.Listen() - Accept() - ServeConn
	//对于http监听器：HandleHTTP http.Serve
	listen, err := net.Listen("tcp", ":1234")
	if err!= nil{
		log.Fatal("listen tcp error: ", err)
	}
	//Accept(): 调用时流程阻塞，直到某个计算机上的某个应用程序与当前程序建立了一个tcp连接，
	//此时Accept()会返回两个值：一个net.Conn类型值，一个error
	//也就是：阻塞等待连接，直到连接上
	conn, err := listen.Accept()
	if err!= nil {
		log.Fatal("Accept error: ", err)
	}

	//单个连接上调用ServeConn管理请求
	rpc.ServeConn(conn)
}
