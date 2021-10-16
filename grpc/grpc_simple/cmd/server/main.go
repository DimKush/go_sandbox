package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	chatpb "github.com/DimKush/go_sandbox/tree/main/grpc/grpc_simple/pkg/api/chatpb"
	"google.golang.org/grpc"
)

const (
	port = ":8082"
)

type server struct {
	chatpb.UnimplementedGreetingSampleServer
}

func(s *server) SayHello(ctx context.Context, in *chatpb.ChatMessage) (*chatpb.ChatAnswer, error) {
	log.Printf("Received %s", in.GetText())
	return &chatpb.ChatAnswer{Answer: fmt.Sprintf("server answered %d", in.GetId())}, nil
}

func main(){
	if  err := run("tcp", port); err != nil {
		fmt.Printf("Critical error during execute run func tcp on port %s Reason: %v", port, err)
		os.Exit(1)
	}
}

func run(network string, port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Cannot create tcp listener on port %s. Reason : %v\n", port, err)
		os.Exit(1)
	}

	grpcNode := grpc.NewServer()
	chatpb.RegisterGreetingSampleServer(grpcNode, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcNode.Serve(listener); err != nil {
		return err
	}
	return nil
}