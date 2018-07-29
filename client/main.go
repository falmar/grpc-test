package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	con, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewTODOClient(con)

	res, err := client.List(context.Background(), &pb.ListRequest{})

	if err != nil {
		log.Fatal(err)
	}

	for {
		todo, err := res.Recv()

		if err != nil {
			break
		}

		fmt.Println(todo)
	}
}
