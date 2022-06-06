package main

import (
	"log"
	"math/rand"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/aditya-todi/grpc-example/ping/proto"
)

func GetClient(addr string) pb.PingServiceClient {
	log.Println("Starting gRPC Go client")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	return pb.NewPingServiceClient(conn)
}

func RandomString() string {
	return strconv.Itoa(rand.Intn(1000000000))
}

func main() {
	conn, err := grpc.Dial("localhost:60967", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c2 := pb.NewPingServiceClient(conn)

	conn, err = grpc.Dial("localhost:60967", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c1 := pb.NewPingServiceClient(conn)

	for i := 0; i < 10; i++ {
		sendPing(c1)
		sendPing(c2)
	}
}
