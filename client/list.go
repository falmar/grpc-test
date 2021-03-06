package main

import (
	"context"
	"fmt"
	"github.com/falmar/grpc-test/pb"
	"io"
)

func List(client pb.TODOClient, flags ListFlags) error {
	res, err := client.List(context.Background(), &pb.ListRequest{
		Query:     flags.query,
		Limit:     flags.limit,
		Offset:    flags.offset,
		Completed: flags.completed,
	})

	if err != nil {
		return err
	}

	var any bool

	for {
		todo, err := res.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		any = true

		fmt.Println(todo)
	}

	if !any {
		fmt.Println("No TODOs found.")
	}

	return nil
}
