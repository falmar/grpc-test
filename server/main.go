package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, _ := net.Listen("tcp", ":8080")

	s := grpc.NewServer()

	pb.RegisterTODOServer(s, &service{})

	s.Serve(lis)
}
