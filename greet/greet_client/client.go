package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client!")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)
	log.Printf("Client is successfully connected: %v", client)

	doUnary(client)
}

func doUnary(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{FirstName: "Sadio", LastName: "Mane"},
	}
	resp, err := client.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while greeting: %v", err)
	}
	log.Printf("Response: %v", resp)

}
