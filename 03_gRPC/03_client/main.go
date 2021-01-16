package main

import (
	"context"
	"fmt"

	"github.com/disharjayanth/Go-modules-crawl-gRPC/03_gRPC/03_client/echo/echo_build"
	"google.golang.org/grpc"
)

func main() {
	// conn is gRPC client connection
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// EchoServerClient is echo client
	echoServerClient := echo_build.NewEchoServerClient(conn)
	// This .Echo method is coming from .proto file (serive method/command)
	ctx := context.Background()
	response, err := echoServerClient.Echo(ctx, &echo_build.EchoRequest{
		Message: "Hello World",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Response from server ->", response.Response)
}
