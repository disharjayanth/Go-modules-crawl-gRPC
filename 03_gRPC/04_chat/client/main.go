package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/disharjayanth/Go-modules-crawl-gRPC/03_gRPC/04_chat/chat/chat_build"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Must have a URL and a client to connect to....")
		return
	}

	ctx := context.Background()

	clientConn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer clientConn.Close()

	chatClient := chat_build.NewChatClient(clientConn)
	clientStream, err := chatClient.Chat(ctx)
	if err != nil {
		panic(err)
	}

	waitC := make(chan struct{})
	go func() {
		for {
			chatMsg, err := clientStream.Recv()
			if err == io.EOF {
				close(waitC)
				return
			} else if err != nil {
				panic(err)
			}
			fmt.Println(chatMsg.User + ":" + chatMsg.Message)
		}
	}()

	fmt.Println("Connection established, type \"quit\" or use ctrl+c to exit.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "quit" {
			err = clientStream.CloseSend()
			if err != nil {
				panic(err)
			}
			break
		}

		err = clientStream.Send(&chat_build.ChatMessage{
			User:    os.Args[1],
			Message: msg,
		})
		if err != nil {
			panic(err)
		}
	}

	<-waitC
}
