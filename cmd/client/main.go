// Package main implements a GRPC client.
package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go-grpc-practice/internal/helloworld"
	"go-grpc-practice/internal/interceptor"
	"log"
	"os"
	"time"

	pb "go-grpc-practice/api/proto/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const name = "world"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print(".env file is absent")
		err = nil
	}
	addr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {
		log.Fatal("SERVER_ADDR env is absent")
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor.ChainClient(helloworld.SignatureSetInterceptor)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("connection closed")
		}
	}(conn)
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}