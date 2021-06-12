package queue_test

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/garupanojisan/mredis/protobuf/pb"
	"github.com/garupanojisan/mredis/queue"
)

func Benchmark_Enqueue(b *testing.B) {
	srv, err := queue.NewQueueServer()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srv.Enqueue(ctx, &pb.EnqueueRequest{
			Key:   "key",
			Value: []byte(uuid.NewV4().String()),
		})
	}
}

func Benchmark_Dequeue(b *testing.B) {
	srv, err := queue.NewQueueServer()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	keys := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = uuid.NewV4().String()
		srv.Enqueue(ctx, &pb.EnqueueRequest{
			Key:   "key",
			Value: []byte(uuid.NewV4().String()),
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srv.Dequeue(ctx, &pb.DequeueRequest{
			Key: "key",
		})
	}
}
