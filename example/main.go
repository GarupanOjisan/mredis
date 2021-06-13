package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"google.golang.org/grpc"

	"github.com/garupanojisan/mredis/protobuf/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// kvs
	kvs := pb.NewKvsClient(conn)
	_, err = kvs.Set(ctx, &pb.SetRequest{
		Key:   "hoge",
		Value: []byte("hoge"),
	})
	if err != nil {
		log.Fatal(err)
	}

	v, err := kvs.Get(ctx, &pb.GetRequest{
		Key: "hoge",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("value = %s\n", string(v.GetValue()))

	// stack
	stack := pb.NewStackClient(conn)
	for i := 0; i < 10; i++ {
		_, err = stack.Push(ctx, &pb.PushRequest{
			Key:   "foo",
			Value: []byte(fmt.Sprintf("%d", i)),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 11; i++ {
		resp, err := stack.Pop(ctx, &pb.PopRequest{
			Key: "foo",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("poped: %v\n", string(resp.Value))
	}

	// queue
	queue := pb.NewQueueClient(conn)
	for i := 0; i < 10; i++ {
		_, err = queue.Enqueue(ctx, &pb.EnqueueRequest{
			Key:   "foo",
			Value: []byte(fmt.Sprintf("%d", i)),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 11; i++ {
		resp, err := queue.Dequeue(ctx, &pb.DequeueRequest{
			Key: "foo",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("dequeued: %v\n", string(resp.Value))
	}

	// common
	common := pb.NewCommonClient(conn)
	r, err := common.Keys(ctx, &pb.KeysRequest{
		Pattern: "f.*",
	})
	if err != nil {
		log.Fatal(err)
	}
	keys := r.GetKeys()
	for i := 0; i < len(keys); i++ {
		fmt.Println(keys[i])
	}

	// sortedset
	sortedset := pb.NewSortedSetClient(conn)
	rankMap := make(map[string]int64)
	for i := 0; i < 10; i++ {
		score := rand.Int63n(100)
		member := fmt.Sprintf("player-%d", i)
		rankMap[member] = score
		_, err := sortedset.ZAdd(ctx, &pb.ZAddRequest{
			Key:    "ranking1",
			Score:  score,
			Member: member,
		})
		fmt.Printf("registered: player-%d, score = %d\n", i, score)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 10; i++ {
		member := fmt.Sprintf("player-%d", i)
		sresp, err := sortedset.ZRank(ctx, &pb.ZRankRequest{
			Key:    "ranking1",
			Member: member,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => rank = %d, score = %d\n", fmt.Sprintf("player-%d", i), sresp.GetRank(), rankMap[member])
	}
}
