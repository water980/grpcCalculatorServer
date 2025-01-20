package main

import (
	"context"
	"grpc2/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.CalculateServer
}

func (s *server) Calculate(ctx context.Context, req *pb.CalculateReq) (*pb.CalculateRes, error) {
	c := int64(req.A) + int64(req.B)
	return &pb.CalculateRes{C: c}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)

	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}
