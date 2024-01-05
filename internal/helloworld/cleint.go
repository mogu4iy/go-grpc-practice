package helloworld

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)
func SignatureSetInterceptor(ctx *context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
	md := metadata.Pairs("signature", "100% guarantee, it's me")
	*ctx = metadata.NewOutgoingContext(*ctx, md)
	return nil
}