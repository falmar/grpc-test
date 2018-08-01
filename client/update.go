package main

import (
	"context"
	"fmt"
	"github.com/falmar/grpc-test/pb"
)

func Update(client pb.TODOClient, flags UpdateFlags) error {
	res, err := client.MarkCompleted(context.Background(), &pb.Todo{
		ID:        flags.id,
		Completed: flags.completed,
	})

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
