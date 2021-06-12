package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/garupanojisan/mredis/common"
	"github.com/garupanojisan/mredis/kvs"
	"github.com/garupanojisan/mredis/protobuf/pb"
	"github.com/garupanojisan/mredis/queue"
	"github.com/garupanojisan/mredis/stack"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	kvs, err := kvs.NewKvsServer()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterKvsServer(srv, kvs)

	stack, err := stack.NewStackServer()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterStackServer(srv, stack)

	queue, err := queue.NewQueueServer()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterQueueServer(srv, queue)

	common, err := common.NewCommonServer()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterCommonServer(srv, common)

	fmt.Println("start listening on :8080")
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
