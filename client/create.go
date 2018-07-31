package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
	"context"
	"fmt"
)

func Create(client pb.TODOClient, flags CreateFlags) error {
	res, err := client.Create(context.Background(), &pb.Todo{
		Name:      flags.name,
		Completed: flags.completed,
	})

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
