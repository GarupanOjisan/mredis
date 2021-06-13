package sortedset

import (
	"context"
	"errors"

	"github.com/garupanojisan/mredis/protobuf/pb"
	"github.com/garupanojisan/mredis/storage"
)

type Server struct {
	pb.UnimplementedSortedSetServer

	memory *storage.Memory
}

func NewSortedSetServer() (*Server, error) {
	return &Server{
		memory: storage.GetInstance(),
	}, nil
}

func (s *Server) ZAdd(ctx context.Context, req *pb.ZAddRequest) (*pb.ZAddResponse, error) {
	if req.GetKey() == "" {
		return nil, errors.New("error invalid key")
	}
	if req.GetMember() == "" {
		return nil, errors.New("error missing member")
	}

	if err := s.memory.ZAdd(ctx, req.GetKey(), req.GetScore(), req.GetMember()); err != nil {
		return nil, err
	}
	return &pb.ZAddResponse{}, nil
}

func (s *Server) ZRank(ctx context.Context, req *pb.ZRankRequest) (*pb.ZRankResponse, error) {
	if req.GetKey() == "" {
		return nil, errors.New("error invalid key")
	}
	if req.GetMember() == "" {
		return nil, errors.New("error missing member")
	}

	rank, err := s.memory.ZRank(ctx, req.GetKey(), req.GetMember())
	if err != nil {
		return nil, err
	}
	return &pb.ZRankResponse{
		Rank: rank,
	}, nil
}
