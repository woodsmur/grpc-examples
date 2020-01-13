package main

import (
	"log"
	"net"

	pb "github.com/woodsmur/grpc-examples/helloworld/internal/proto/helloworld"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":50051"

// server is used to implement pb.GreeterServer.
type server struct{}

// SayHello implements pb.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
