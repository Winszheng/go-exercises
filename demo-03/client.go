package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	//1.和grpc服务建立连接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err!= nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//2.基于已经建立的连接构造HelloServiceClient对象
	//这样返回的是一个接口对象，这样就可以调用服务器提供的grpc方法
	client := NewHelloServiceClient(conn)

	//3.调用grpc方法
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
