package kvs_test

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/garupanojisan/mredis/kvs"
	"github.com/garupanojisan/mredis/protobuf/pb"
)

func Benchmark_Set(b *testing.B) {
	srv, err := kvs.NewKvsServer()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srv.Set(ctx, &pb.SetRequest{
			Key:   uuid.NewV4().String(),
			Value: []byte(uuid.NewV4().String()),
		})
	}
}

func Benchmark_Get(b *testing.B) {
	srv, err := kvs.NewKvsServer()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	keys := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = uuid.NewV4().String()
		srv.Set(ctx, &pb.SetRequest{
			Key:   uuid.NewV4().String(),
			Value: []byte(uuid.NewV4().String()),
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := keys[i]
		srv.Get(ctx, &pb.GetRequest{
			Key: key,
		})
	}
}
