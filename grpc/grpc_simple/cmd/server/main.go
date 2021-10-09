package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	""
)

const (
	port = ":80821"
)

func main(){

}

type server struct {
	pb.UnimplementedGreetingSampleServer
}

func run(network string, port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Cannot create tcp listener on port %s. Reason : %v\n", port, err)
		os.Exit(1)
	}

	grpcNode := grpc.NewServer()
	return nil
}