package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	unit "github.com/DimKush/go_sandbox/tree/main/grpc/grpc_simple/pkg/api/unit"
	"google.golang.org/grpc"
)

const (
	port = ":8082"
)

type server struct {
	unit.UnimplementedGreetingSampleServer
}

func (s *server) GetUnitById(ctx context.Context, in *unit.UnitId) (*unit.UnitsResponce, error) {
	id := in.GetUnitId()

}

func main() {
	if err := run("tcp", port); err != nil {
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
	unit.RegisterGreetingSampleServer(grpcNode, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcNode.Serve(listener); err != nil {
		return err
	}
	return nil
}
