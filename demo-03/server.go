package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listen, err := net.Listen("tcp", ":1234")
	if err!= nil {
		log.Fatal(err)
	}

	grpcServer.Serve(listen)
}
