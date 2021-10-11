package main

import (
	"google.golang.org/grpc"
	chatpb "github.com/DimKush/go_sandbox/tree/main/grpc/grpc_simple/pkg/chatpb"
)

const (
	host = "localhost"
	port = ":8082"
	full_address = host + port
)

func main() {

}

func runClient() error{
	connect, err := grpc.Dial(full_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		return err
	}

	defer connect.Close()

	clnt := chatpb.NewGreetingSampleClient()
}