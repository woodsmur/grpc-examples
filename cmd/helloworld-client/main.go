package main

import (
	"flag"
	"log"

	pb "github.com/woodsmur/grpc-examples/helloworld/internal/proto/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	addr := flag.String("addr", address, "server address")
	name := flag.String("name", defaultName, "hello who?")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)

	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}
	log.Printf("Greeting: %s", r.Message)

}
