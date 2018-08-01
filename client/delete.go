package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
	"context"
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
