package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/woodsmur/grpc-examples/helloworld/internal/proto/helloworld"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.String("port", ":50051", "Listen port")
)

// server is used to implement pb.GreeterServer.
type server struct{}

// SayHello implements pb.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if *certFile != "" && *keyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(s, &server{})
	server := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, server)
	server.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	s.Serve(lis)
}
