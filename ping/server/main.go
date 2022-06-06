package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"

	pb "github.com/aditya-todi/grpc-example/ping/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

type server struct {
	pb.UnimplementedPingServiceServer
}

func unaryLogInterceptor(logger *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		e := logger.Info()
		e.Msg("Logging through interceptor")
		e.Str("method", info.FullMethod)
		return handler(ctx, req)
	}
}

func RandomString() string {
	return strconv.Itoa(rand.Intn(1000000000))
}

func NewLogger(logsFilePath string, isDebug bool) *zerolog.Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.TraceLevel
	}

	logsFile, err := os.OpenFile(logsFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		print(err.Error())
		panic("cannot open logs file")
	}

	zerolog.SetGlobalLevel(logLevel)
	l := zerolog.New(logsFile).With().Timestamp().Logger()
	return &l
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger := NewLogger("/grpc-example/logs/"+RandomString()+".log", true)

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryLogInterceptor(logger)))

	pb.RegisterPingServiceServer(s, &server{})

	e := logger.Info()
	e.Msg("server listening at " + lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
