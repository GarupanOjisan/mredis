package stack_test

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/garupanojisan/mredis/protobuf/pb"
	"github.com/garupanojisan/mredis/stack"
)

func Benchmark_Push(b *testing.B) {
	srv, err := stack.NewStackServer()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srv.Push(ctx, &pb.PushRequest{
			Key:   "key",
			Value: []byte(uuid.NewV4().String()),
		})
	}
}

func Benchmark_Pop(b *testing.B) {
	srv, err := stack.NewStackServer()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	keys := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = uuid.NewV4().String()
		srv.Push(ctx, &pb.PushRequest{
			Key:   "key",
			Value: []byte(uuid.NewV4().String()),
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srv.Pop(ctx, &pb.PopRequest{
			Key: "key",
		})
	}
}
