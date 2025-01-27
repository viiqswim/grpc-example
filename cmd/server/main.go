package main

import (
	"context"
	"log"
	"net"

	greeter "grpc_example/pkg/greeter/api/proto/v1"

	"google.golang.org/grpc"
)

type greeterServer struct {
	greeter.UnimplementedGreeterServiceServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{Message: "Hello, " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServiceServer(grpcServer, &greeterServer{})
	log.Println("Server is running on port 50051...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
