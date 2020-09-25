package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	v1 "github.com/Winszheng/go-grpc-http-rest-microservice/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"time"
)

const(
	apiVersion = "v1"
)

func main()  {
	certificate, err := tls.LoadX509KeyPair("../../conf/client-keys/client.crt", "../../conf/client-keys/client.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../conf/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{certificate},
		ServerName:         "server.grpc.io", // NOTE: this is required!
		RootCAs:            certPool,
	})

/*
	//構造客戶端使用的證書對象
	creds, err := credentials.NewClientTLSFromFile("../../conf/keys/server.crt","server.grpc.io")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}

 */
	//get configuration
	address := flag.String("server", "", "grpc server in format host:port")
	flag.Parse()

	//建立到server的连接
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(creds))
	if err!= nil {
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()

	c := v1.NewToDoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	prefix := t.Format(time.RFC3339Nano)

	//CALL create
	req1 := v1.CreateRequest{
		Api:  apiVersion,
		ToDo: &v1.ToDo{
			Title:       "title("+prefix+")",
			Description: "desc("+prefix+")",
			Reminder:    reminder,
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalln("Create failed: ", err)
	}
	log.Printf("create result: %v\n", res1)
	id := res1.Id

	//read
	req2 := v1.ReadRequest{
		Api: apiVersion,
		Id:  id,
	}
	res2, err := c.Read(ctx, &req2)
	if err!= nil {
		log.Fatalf("read failed: %v", err)
	}
	log.Println("read result: ", res2)

	//update
	req3 := v1.UpdateRequest{
		Api:  apiVersion,
		ToDo: &v1.ToDo{
			Id:          res2.ToDo.Id,
			Title:       res2.ToDo.Title,
			Description: res2.ToDo.Description+" updated",
			Reminder:    res2.ToDo.Reminder,
		},
	}
	res3, err := c.Update(ctx, &req3)
	if err != nil {
		log.Fatalf("update failed: %v", err)
	}
	log.Printf("update result: %+v\n", res3)

	//readall
	req4 := v1.ReadAllRequest{Api:apiVersion}
	res4, err := c.ReadAll(ctx, &req4)
	if err!= nil {
		log.Fatalf("readall failed: %v", err)
	}
	log.Println("readall result: %v\n", res4)

	//delete
	req5 := v1.DeleteRequest{
		Api: apiVersion,
		Id:  id,
	}
	res5, err := c.Delete(ctx, &req5)
	if err!= nil {
		log.Fatalf("delete failed: %v", err)
	}
	log.Println("delete result: ", res5)
}
