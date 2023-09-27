package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "/Echo-project-gRPC-with-go/echo/Echo/Echo"

	"google.golang.org/grpc"
)

var defaultPort int = 7000
var (
	port = flag.Int("port", defaultPort, "Server Port")
)

type server struct {
	pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("The message is : %v", in.GetEchoRequest())
	var message string = in.GetEchoRequest()
	return &pb.EchoResponse{EchoResponse: message}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterEchoServiceServer(s, &server{})

	log.Printf("Server is listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
