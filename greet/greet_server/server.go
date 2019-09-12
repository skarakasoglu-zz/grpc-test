package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"io"
	"log"
	"net"
	"strconv"
)

type server struct{}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Println("Client invoked GreetManyTimes func")
	result := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// other implementation
			resp := &greetpb.LongGreetResponse{
				Result: result,
			}
			return stream.SendAndClose(resp)
		}
		if err != nil {
			log.Fatalf("failed while receiving: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}

}

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Println("Client invoked GreetManyTimes func")

	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: "Hello " + firstName + " number " + strconv.Itoa(i),
		}

		stream.Send(res)
	}

	return nil
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.Greeting.GetFirstName()
	result := "Hello " + firstName
	response := &greetpb.GreetResponse{Result: result}

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
