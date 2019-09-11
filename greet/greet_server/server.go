package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"log"
	"net"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.Greeting.GetFirstName()
	result := "Hello " + firstName
	response := &greetpb.GreetResponse{Result:result}

	return response, nil
}

func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Error occured while listening: %v", err)
	}

	serv := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(serv, &server{})

	if err = serv.Serve(lis); err != nil {
		log.Fatalf("Error occured while serving: %v", err)
	}
}
