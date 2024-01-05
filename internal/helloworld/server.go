package helloworld

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"

	pb "go-grpc-practice/api/proto/greeter"
)
type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func SignatureCheckInterceptor(ctx context.Context, _ any) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}
	sign, ok := md["signature"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Signature token is not supplied")
	}
	log.Printf("signature: %v", sign)
	return nil, nil
}