package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
	"context"
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
		Query:     rq.Query,
		Completed: rq.Completed,
		Limit:     rq.Limit,
		Offset:    rq.Offset,
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

func (s *service) Create(ctx context.Context, rq *pb.Todo) (*pb.Todo, error) {
	t := &TODO{
		Name:      rq.Name,
		Completed: rq.Completed,
	}

	err := s.dl.Create(t)

	if err != nil {
		return nil, err
	}

	return &pb.Todo{
		ID:        t.ID,
		Name:      t.Name,
		Completed: t.Completed,
	}, nil
}

func (s *service) MarkCompleted(ctx context.Context, rq *pb.Todo) (*pb.Todo, error) {
	t := &TODO{
		ID:        rq.ID,
		Completed: rq.Completed,
	}

	err := s.dl.Update(t)

	if err != nil {
		return nil, err
	}

	return &pb.Todo{
		ID:        t.ID,
		Name:      t.Name,
		Completed: t.Completed,
	}, nil
}

func (s *service) Delete(ctx context.Context, rq *pb.Todo) (*pb.Empty, error) {
	return &pb.Empty{}, s.dl.Delete(rq.ID)
}
