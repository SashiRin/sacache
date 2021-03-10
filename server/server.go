package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/sashirin/sacache"
	pb "github.com/sashirin/sacache/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	version = "0.1.0"
)

var (
	port             int
	logfile          string
	cacheServiceName string

	cache *sacache.SaCache
)

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

	grpcServer := grpc.NewServer()

	pb.RegisterCacheServiceServer(grpcServer, sacache.NewSaCache(cacheServiceName))

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
