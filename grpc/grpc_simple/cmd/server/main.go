package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	chatpb "github.com/DimKush/go_sandbox/tree/main/grpc/grpc_simple/pkg/chatpb"
	server "github.com/DimKush/go_sandbox/tree/main/grpc/grpc_simple/server"
)


type server struct {
	chatpb.UnimplementedGreetingSampleServer
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