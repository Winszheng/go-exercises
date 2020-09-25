#### environment

ubuntu 20.04

If ur os is windows, u should modify the code.

#### how to run

1.建立相应数据库表

2.运行服务端

```g
cd cmd/server
go build .
server.exe -grpc-port=9090 -db-host=<HOST>:3306 -db-user=<USER> -db-password=<PASSWORD> -db-schema=<SCHEMA>
```

for example:

```
cd cmd/server
go build .
./server -grpc-port=9090 -db-host=localhost:3306 -db-user=root -db-password=2333 -db-schema=todo

```

3.运行客户端

```
cd cmd/client-grpc
go build .
./client-grpc -server=localhost:9090
```

