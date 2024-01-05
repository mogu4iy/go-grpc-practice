// Package main implements a GRPC server.
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-grpc-practice/internal/helloworld"
	"go-grpc-practice/internal/interceptor"
	"log"
	"net"
	"os"
	"strconv"

	pb "go-grpc-practice/api/proto/greeter"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print(".env file is absent")
		err = nil
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT env is absent")
	}
	_, err = strconv.Atoi(port)
	if err != nil{
		log.Fatal("port is not number")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	hw := &helloworld.Server{}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ChainServer(helloworld.SignatureCheckInterceptor)),
	)
	pb.RegisterGreeterServer(s, hw)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Printf("server listening at %v", lis.Addr())
	}
}