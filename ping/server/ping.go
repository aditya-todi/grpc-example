package main

import (
	"context"
	"log"

	pb "github.com/aditya-todi/grpc-example/ping/proto"
)

func (s *server) Ping(ctx context.Context, in *pb.PingMessage) (*pb.PongMessage, error) {
	log.Printf("Ping rpc invoked with message %v", in.Message)
	return &pb.PongMessage{
		Message: "Responding to message: " + in.Message,
	}, nil
}
