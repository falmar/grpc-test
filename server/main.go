package main

import (
	"github.com/falmar/grpc-test/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, _ := net.Listen("tcp", ":8080")

	s := grpc.NewServer()

	dl := NewDataLayer()
	svc := NewService(dl)

	pb.RegisterTODOServer(s, svc)

	s.Serve(lis)
}
