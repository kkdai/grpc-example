package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/kkdai/grpc-example/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayStreamHello(in *pb.HelloRequest, stream pb.Greeter_SayStreamHelloServer) error {
	respones := []string{"hello", "good", "morning"}
	for _, res := range respones {
		stream.Send(&pb.HelloReply{Message: res + in.Name})
		time.Sleep(2 * time.Second)
	}
	return nil
}

func main() {
	srv := new(server)
	c, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, srv)
	s.Serve(c)
}
