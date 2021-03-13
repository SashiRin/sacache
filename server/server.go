package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/sashirin/sacache"
	pb "github.com/sashirin/sacache/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	version = "0.2.0"
)

var (
	port             int
	logfile          string
	cacheServiceName string

	cache *sacache.SaCache
)

// CacheServer is the cache service server.
type CacheServer struct {
	pb.UnimplementedCacheServiceServer
}

// Get returns the CacheItem pointer of given key.
func (s *CacheServer) Get(ctx context.Context, args *pb.GetKey) (*pb.CacheItem, error) {
	key := args.Key
	item, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	log.Printf("get item with key: %v", key)
	return &pb.CacheItem{
		Key:        key,
		Value:      item.Value(),
		ExpireTime: item.ExpireTime().Format(time.RFC3339),
	}, nil
}

// Set add new k-v pair in the cache.
func (s *CacheServer) Set(ctx context.Context, item *pb.CacheItem) (*pb.Success, error) {
	expire, _ := time.Parse(time.RFC3339, item.ExpireTime)
	err := cache.Set(item.Key, item.Value, expire)
	if err != nil {
		return &pb.Success{
			Success: false,
		}, err
	}
	log.Printf("set item: %v %v %v", item.Key, item.Value, expire)
	return &pb.Success{
		Success: true,
	}, nil
}

// Delete deletes value given key.
func (s *CacheServer) Delete(ctx context.Context, args *pb.GetKey) (*pb.Success, error) {
	key := args.Key
	err := cache.Delete(key)
	if err != nil {
		return &pb.Success{
			Success: false,
		}, err
	}
	log.Printf("delete item with key: %v", key)
	return &pb.Success{
		Success: true,
	}, nil
}

func init() {
	flag.IntVar(&port, "port", 9999, "the port to listen.")
	flag.StringVar(&logfile, "logfile", "", "path of logfile.")
	flag.StringVar(&cacheServiceName, "name", "SaCacheServiceName", "name of service.")
}

func main() {
	flag.Parse()

	fmt.Printf("SaCache Server v%s\n", version)

	var logger *log.Logger

	if logfile == "" {
		logger = log.New(os.Stdout, "", log.LstdFlags)

	} else {
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		logger = log.New(f, "", log.LstdFlags)
	}

	cache = sacache.NewSaCache(cacheServiceName)

	grpcServer := grpc.NewServer()

	pb.RegisterCacheServiceServer(grpcServer, &CacheServer{})

	reflection.Register(grpcServer)

	logger.Print("cache initialized successfully.")

	address := ":" + strconv.Itoa(port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	logger.Printf("starting server on %v", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("err in serving gRPC %v\n", err)
	}
}
