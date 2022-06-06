package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/aditya-todi/grpc-example/ping/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

type server struct {
	pb.UnimplementedPingServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
