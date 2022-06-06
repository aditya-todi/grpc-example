package main

import (
	"context"
	"log"

	pb "github.com/aditya-todi/grpc-example/ping/proto"
)

func sendPing(c pb.PingServiceClient) {
	log.Println("Ping invoked from client")
	res, err := c.Ping(context.Background(), &pb.PingMessage{
		Message: "Hello server, this is client",
	})

	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}

	log.Printf("Response: %s\n", res.Message)
}
