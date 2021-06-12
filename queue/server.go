package queue

import (
	"context"
	"errors"

	"github.com/garupanojisan/mredis/protobuf/pb"
	"github.com/garupanojisan/mredis/storage"
)

type Server struct {
	pb.UnimplementedQueueServer

	memory *storage.Memory
}

func NewQueueServer() (*Server, error) {
	return &Server{
		memory: storage.GetInstance(),
	}, nil
}

func (s *Server) Enqueue(ctx context.Context, req *pb.EnqueueRequest) (*pb.EnqueueResponse, error) {
	if req.GetKey() == "" {
		return nil, errors.New("error invalid key")
	}
	if req.GetValue() == nil {
		return nil, errors.New("error empty value")
	}

	err := s.memory.Enqueue(ctx, req.GetKey(), req.GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.EnqueueResponse{}, nil
}

func (s *Server) Dequeue(ctx context.Context, req *pb.DequeueRequest) (*pb.DequeueResponse, error) {
	if req.GetKey() == "" {
		return nil, errors.New("error invalid key")
	}

	v, err := s.memory.Dequeue(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.DequeueResponse{
		Value: v,
	}, nil
}
