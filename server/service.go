package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
)

func NewService(dl DataLayer) *service {
	return &service{
		dl: dl,
	}
}

type service struct {
	dl DataLayer
}

func (s *service) List(rq *pb.ListRequest, stm pb.TODO_ListServer) error {

	todos, err := s.dl.List(ListFilter{
		rq.Query,
		false,
		rq.Limit,
		rq.Offset,
	})

	if err != nil {
		return err
	}

	for _, v := range todos {
		stm.Send(&pb.Todo{
			ID:        v.ID,
			Name:      v.Name,
			Completed: v.Completed,
		})
	}

	return nil
}
