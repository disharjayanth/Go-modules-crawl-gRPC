package main

import (
	"context"
	"fmt"
	"net"

	"github.com/disharjayanth/Go-modules-crawl-gRPC/03_gRPC/02_server/echo/echo_op"
	"google.golang.org/grpc"
)

// EchoServer implements UnimplementedEchoServerServer struct and Echo method attached to it
type EchoServer struct {
	echo_op.UnimplementedEchoServerServer
}

// Echo method recevies EchoRequest and returns EchoResponse
func (e *EchoServer) Echo(ctx context.Context, req *echo_op.EchoRequest) (*echo_op.EchoResponse, error) {
	return &echo_op.EchoResponse{
		Response: "My Echo from server: " + req.Message,
	}, nil
}

func main() {
	TPCListener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	gRPCServer := grpc.NewServer()

	echoServer := &EchoServer{}

	echo_op.RegisterEchoServerServer(gRPCServer, echoServer)

	fmt.Println("Now serving @:3000")
	err = gRPCServer.Serve(TPCListener)
	if err != nil {
		panic(err)
	}
}
