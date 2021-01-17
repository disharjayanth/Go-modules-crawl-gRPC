package main

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/disharjayanth/Go-modules-crawl-gRPC/03_gRPC/04_chat/chat/chat_build"
	"google.golang.org/grpc"
)

type Connection struct {
	conn chat_build.Chat_ChatServer
	send chan *chat_build.ChatMessage
	quit chan struct{}
}

func NewConnection(conn chat_build.Chat_ChatServer) *Connection {
	c := &Connection{
		conn: conn,
		send: make(chan *chat_build.ChatMessage),
		quit: make(chan struct{}),
	}

	go c.start()
	return c
}

func (c *Connection) Close() error {
	close(c.send)
	close(c.quit)
	return nil
}

func (c *Connection) Send(msg *chat_build.ChatMessage) {
	defer func() {
		// Ignore any error about sending into closed channel.
		recover()
	}()
	c.send <- msg
}

func (c *Connection) start() {
	running := true
	for running {
		select {
		case msg := <-c.send:
			c.conn.Send(msg)
		case <-c.quit:
			running = false
		}
	}
}

func (c *Connection) GetMessage(broadcast chan<- *chat_build.ChatMessage) error {
	for {
		msg, err := c.conn.Recv()
		if err == io.EOF {
			c.Close()
			return nil
		} else if err != nil {
			c.Close()
			return err
		}
		go func(msg *chat_build.ChatMessage) {
			select {
			case broadcast <- msg:
			case <-c.quit:
			}
		}(msg)
	}
}

type ChatServer struct {
	broadcast   chan *chat_build.ChatMessage
	quit        chan struct{}
	connections []*Connection
	connLock    sync.Mutex
	chat_build.UnimplementedChatServer
}

func NewChatServer() *ChatServer {
	srv := &ChatServer{
		broadcast: make(chan *chat_build.ChatMessage),
		quit:      make(chan struct{}),
	}
	go srv.start()
	return srv
}

func (c *ChatServer) Close() error {
	close(c.quit)
	return nil
}

func (c *ChatServer) start() {
	running := true
	for running {
		select {
		case msg := <-c.broadcast:
			c.connLock.Lock()
			for _, v := range c.connections {
				go v.Send(msg)
			}
			c.connLock.Unlock()
		case <-c.quit:
			running = false
		}
	}
}

func (c *ChatServer) Chat(stream chat_build.Chat_ChatServer) error {
	conn := NewConnection(stream)

	c.connLock.Lock()
	c.connections = append(c.connections, conn)
	c.connLock.Unlock()

	err := conn.GetMessage(c.broadcast)

	c.connLock.Lock()
	for i, v := range c.connections {
		if v == conn {
			c.connections = append(c.connections[:i], c.connections[i+1:]...)
		}
	}
	c.connLock.Unlock()

	return err
}

func main() {
	lst, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	srv := NewChatServer()

	chat_build.RegisterChatServer(s, srv)

	fmt.Println("Listening PORT @:3000")
	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}
}
