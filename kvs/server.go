package kvs

import (
	"context"
	"errors"

	"github.com/garupanojisan/mredis/protobuf/pb"
	"github.com/garupanojisan/mredis/storage"
)

type Server struct {
	pb.UnimplementedKvsServer

	memory *storage.Memory
}

func NewKvsServer() (*Server, error) {
	return &Server{
		memory: storage.GetInstance(),
	}, nil
}

func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	if req.GetKey() == "" {
		return nil, errors.New("error invalid key")
	}
	if req.GetValue() == nil {
		return nil, errors.New("error empty value")
	}

	if err := s.memory.Set(ctx, req.GetKey(), req.GetValue()); err != nil {
		return nil, err
	}
	return &pb.SetResponse{}, nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if req.GetKey() == "" {
		return nil, errors.New("error invalid key")
	}

	v, err := s.memory.Get(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{
		Value: v,
	}, nil
}
