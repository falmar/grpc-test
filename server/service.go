package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
	"strings"
)

type TODOItem map[string]interface{}

var todoRepo = []TODOItem{
	{
		"ID":        "uuid-1",
		"Name":      "Make a cli client",
		"Completed": false,
	},
	{
		"ID":        "uuid-2",
		"Name":      "Implement list",
		"Completed": false,
	}, {
		"ID":        "uuid-3",
		"Name":      "implement get",
		"Completed": false,
	},
}

type service struct{}

func (s *service) List(rq *pb.ListRequest, stm pb.TODO_ListServer) error {
	var filtered []TODOItem

	for _, v := range todoRepo {
		if !strings.Contains(v["Name"].(string), rq.Query) {
			continue
		}

		filtered = append(filtered, v)
	}

	max := int64(len(filtered))

	if rq.Limit < 0 {
		rq.Limit = 0
	} else if rq.Limit > max {
		rq.Limit = max
	}

	if rq.Offset < 0 {
		rq.Offset = 0
	} else if rq.Offset+1 > max {
		rq.Offset = max
	}

	for _, v := range filtered[rq.Offset:rq.Limit] {
		stm.Send(&pb.Todo{
			ID:        v["ID"].(string),
			Name:      v["Name"].(string),
			Completed: v["Completed"].(bool),
		})
	}

	return nil
}
