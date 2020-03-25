package main

import (
	"flag"
	"log"
	"os"

	"google.golang.org/grpc/credentials"

	pb "github.com/woodsmur/grpc-examples/helloworld/internal/proto/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	name       = flag.String("name", "world", "hello who?")
	addr       = flag.String("addr", "localhost:50051", "server address")
	serverName = flag.String("server_name", "helloworld.grpc.example.com", "Server name override")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	if addrEnv := os.Getenv("SERVER_ADDR"); addrEnv != "" {
		log.Printf("use SERVER_ADDR : %v", addrEnv)
		*addr = addrEnv
	}
	log.Printf("server addr : %v", *addr)

	var conn *grpc.ClientConn
	var err error
	if *certFile != "" {
		creds, _ := credentials.NewClientTLSFromFile(*certFile, *serverName)
		conn, err = grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalf("did not connect: %v with cert file : %v", err, *certFile)
		}
	} else {
		conn, err = grpc.Dial(*addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
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
