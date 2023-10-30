package main

import (
	"context"
	"fmt"
	"log"

	"github.com/longpt233/go-grpc/external/proto"

	"google.golang.org/grpc"
)

func main() {
	// thiết lập kết nối với gRPC service
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// xây dựng đối tượng `HelloServiceClient` dựa trên kết nối đã thiết lập
	client := proto.NewHelloServiceClient(conn)
	reply, err := client.SayHello(context.Background(), &proto.HelloMessage{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
