package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	pb "github.com/kkdai/grpc-example/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50058"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("Start to SayHello")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloStreamServer(in *pb.HelloRequest, stream pb.Greeter_SayHelloStreamServerServer) error {
	log.Println("Start to SayHellpStreamServer")
	respones := []string{"hello", "good", "morning"}
	for _, res := range respones {
		stream.Send(&pb.HelloReply{Message: res + in.Name})
		time.Sleep(2 * time.Second)
	}
	return nil
}

func (s *server) SayHelloStreamClient(stream pb.Greeter_SayHelloStreamClientServer) error {
	var totalNames string
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.GetName())
		totalNames = totalNames + " " + res.GetName()
	}

	return stream.SendAndClose(&pb.HelloReply{Message: "Hi " + totalNames})
}

func (s *server) FileProcess(ctx context.Context, in *pb.FileRequest) (*pb.FileReply, error) {
	log.Println("Start to receive file...")
	log.Println("Totoal size:", len(in.GetImage()))
	return &pb.FileReply{Image: in.GetImage(), Result: "Succeed"}, nil
}

func main() {
	srv := new(server)
	c, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.MaxMsgSize(1024 * 1024 * 8),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(s, srv)
	log.Println("Server starting...")
	s.Serve(c)
}
