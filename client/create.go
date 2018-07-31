package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
	"context"
	"encoding/json"
	"fmt"
)

func Create(client pb.TODOClient, flags CreateFlags) error {
	var data struct {
		Name      string
		Completed bool
	}

	err := json.Unmarshal([]byte(flags.data), &data)

	if err != nil {
		return err
	}

	res, err := client.Create(context.Background(), &pb.Todo{
		Name:      data.Name,
		Completed: data.Completed,
	})

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
