package main

import (
	"context"
	pb "/Echo-project-gRPC-with-go/echo/Echo/Echo"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultValue = "Hello"
)

var (
	addr        = flag.String("addr", "localhost:7000", "Server address")
	echoRequest = flag.String("echo", defaultValue, "echo message ")
)

func main() {
	flag.Parse()
	con, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection not Establish %v", err)
	}else{
		log.Println("Connection Established")
	}
	c := pb.NewEchoServiceClient(con)
	defer con.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoRequest{EchoRequest: *echoRequest})

	if err != nil {
		log.Fatalf("Echo message return an error %v", err)
	}

	log.Println("Echo : ",r.GetEchoResponse())
	

}
