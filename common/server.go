package common

import (
	"context"
	"errors"

	"github.com/garupanojisan/mredis/protobuf/pb"
	"github.com/garupanojisan/mredis/storage"
)

type Server struct {
	pb.UnimplementedCommonServer

	memory *storage.Memory
}

func NewCommonServer() (*Server, error) {
	return &Server{
		memory: storage.GetInstance(),
	}, nil
}

func (s *Server) Keys(ctx context.Context, req *pb.KeysRequest) (*pb.KeysResponse, error) {
	if req.GetPattern() == "" {
		return nil, errors.New("error missing pattern")
	}

	keys, err := s.memory.Keys(ctx, req.GetPattern())
	if err != nil {
		return nil, err
	}
	return &pb.KeysResponse{
		Keys: keys,
	}, nil
}
