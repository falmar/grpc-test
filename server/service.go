package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
)

var todoRepo = []map[string]interface{}{
	map[string]interface{}{
		"ID":        "uuid-1",
		"Name":      "Make a cli client",
		"Completed": false,
	},
	map[string]interface{}{
		"ID":        "uuid-2",
		"Name":      "Implement list",
		"Completed": false,
	},
	map[string]interface{}{
		"ID":        "uuid-3",
		"Name":      "implement get",
		"Completed": false,
	},
}

type service struct{}

func (s *service) List(rq *pb.ListRequest, stm pb.TODO_ListServer) error {

	for _, v := range todoRepo {
		stm.Send(&pb.Todo{
			ID:        v["ID"].(string),
			Name:      v["Name"].(string),
			Completed: v["Completed"].(bool),
		})
	}

	return nil
}
