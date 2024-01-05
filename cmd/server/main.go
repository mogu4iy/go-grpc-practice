// Package main implements a GRPC server.
package main

import (
	"flag"
	"fmt"
	"go-grpc-practice/internal/helloworld"
	"go-grpc-practice/internal/interceptor"
	"log"
	"net"

	pb "go-grpc-practice/api/proto/greeter"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)


func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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