package main

import (
	"context"
	"github.com/falmar/grpc-test/pb"
)

func Delete(client pb.TODOClient, flags DeleteFlags) error {
	_, err := client.Delete(context.Background(), &pb.Todo{
		ID: flags.id,
	})

	if err != nil {
		return err
	}

	return nil
}
