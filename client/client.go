package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/sashirin/sacache/proto"
	"google.golang.org/grpc"
)

const (
	address            = ":9999"
	defaultKey         = "world"
	defaultValue       = "23333"
	defaultDurationStr = "200s"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCacheServiceClient(conn)

	// Contact the server and print out its response.
	key := defaultKey
	value := defaultValue
	durationStr := defaultDurationStr
	if len(os.Args) > 1 {
		key = os.Args[1]
		value = os.Args[2]
		durationStr = os.Args[3]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	duration, _ := time.ParseDuration(durationStr)
	expire := time.Now().Add(duration).Format(time.RFC3339)

	r, err := c.Set(ctx, &pb.CacheItem{Key: key, Value: value, ExpireTime: expire})
	if err != nil || !r.GetSuccess() {
		log.Fatalf("could not set item: %v %v %v", key, value, expire)
	}
	log.Printf("Set item %v %v %v successful", key, value, expire)
	getKey := &pb.GetKey{
		Key: key,
	}
	item, _ := c.Get(ctx, getKey)
	if err != nil {
		log.Fatalf("could not found key: %v", err)
	}
	if item.GetKey() != key {
		log.Fatalf("key error: %v", key)
	}
}
