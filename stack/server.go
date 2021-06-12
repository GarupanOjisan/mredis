package stack

import (
	"context"
	"errors"

	"github.com/garupanojisan/mredis/protobuf/pb"
	"github.com/garupanojisan/mredis/storage"
)

type Server struct {
	pb.UnimplementedStackServer

	memory *storage.Memory
}

func NewStackServer() (*Server, error) {
	return &Server{
		memory: storage.GetInstance(),
	}, nil
}

func (s *Server) Push(ctx context.Context, req *pb.PushRequest) (*pb.PushResponse, error) {
	if req.GetKey() == "" {
		return nil, errors.New("error invalid key")
	}
	if req.GetValue() == nil {
		return nil, errors.New("error empty value")
	}

	err := s.memory.Push(ctx, req.GetKey(), req.GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.PushResponse{}, nil
}

func (s *Server) Pop(ctx context.Context, req *pb.PopRequest) (*pb.PopResponse, error) {
	if req.GetKey() == "" {
		return nil, errors.New("error invalid key")
	}

	v, err := s.memory.Pop(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.PopResponse{
		Value: v,
	}, nil
}
