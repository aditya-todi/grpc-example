FROM golang:alpine

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev protoc

RUN mkdir /grpc-example 

WORKDIR /grpc-example

COPY go.mod /grpc-example
COPY go.sum /grpc-example
COPY ping /grpc-example/ping

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN protoc --go_opt=module=github.com/aditya-todi/grpc-example --go_out=. --go-grpc_opt=module=github.com/aditya-todi/grpc-example --go-grpc_out=. ./ping/proto/request.proto
RUN go build -o build/ping/server ./ping/server
# RUN go build -o build/ping/client ./ping/client

CMD ["./build/ping/server"]
EXPOSE 9000/tcp