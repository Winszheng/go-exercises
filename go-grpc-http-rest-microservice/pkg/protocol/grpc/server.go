package grpc

import (
	"context"
	v1 "github.com/Winszheng/go-grpc-http-rest-microservice/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
	"os/signal"
)

//RunServer函数负责注册ToDo服务以及启动grpc服务
func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	//1.listen函数创建服务端, 指定监听端口
	listen, err := net.Listen("tcp", ":"+port)
	if err!= nil {
		return err
	}
	//2.从输入证书文件和密钥文件为服务端构造tls凭证
	//注意这里使用的路径，是相对可执行文件cmd/server/main.go去搜索的 --- 迷惑行为
	creds, err := credentials.NewServerTLSFromFile("../../conf/keys/server.crt","../../conf/keys/server.key")

	if err != nil {
		log.Fatalf("failed to generate credentials : %v", err)
	}
	//3.创建grpc服务器的一个实例, 并开启tls认证
	server := grpc.NewServer(grpc.Creds(creds)) //grpc.Creds(creds)将证书包装为选项
	//4.在grpc服务器(server)注册项目的服务实现(v1API)
	v1.RegisterToDoServiceServer(server, v1API) //注册服务器

	//graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for range c {
			log.Println("shutting down grpc server")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	//start grpc server
	log.Println("starting grpc server...")
	//5.用服务器Serve()方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者stop()被调用
	return server.Serve(listen)
}
