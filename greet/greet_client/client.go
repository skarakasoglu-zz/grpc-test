package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"io"
	"log"
	"time"
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

	// doUnary(client)
	// doServerStreaming(client)
	// doClientStreaming(client)
	doBiDiStreaming(client)
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

func doServerStreaming(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a server streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Roberto",
			LastName:  "Firmino",
		},
	}

	resultStream, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error call to greet many times: %v", err)
	}

	for {
		val, err := resultStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive from stream: %v", err)
		}
		log.Printf("Server response: %v", val)
	}
}

func doClientStreaming(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a client streaming RPC...")

	people := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{FirstName: "Roberto"}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{FirstName: "Sadio"}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{FirstName: "Wijnaldum"}},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("failed while calling LongGreet: %v", err)
	}

	for _, val := range people {

		err := stream.Send(val)
		if err != nil {
			log.Fatalf("failed while sending to the stream: %v", err)
		}

		log.Printf("Sending %v", val.GetGreeting().GetFirstName())
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("internal server error: %v", err)
	}

	log.Println(resp)

}

func doBiDiStreaming(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a bidi streaming RPC...")

	people := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{FirstName: "Roberto"}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{FirstName: "Sadio"}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{FirstName: "Wijnaldum"}},
	}

	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("failed to call GreetEveryone: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range people {
			log.Printf("Sending: %v\n", req)
			err := stream.Send(req)
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatalf("failed to send to stream: %v", err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatalf("failed to receive from stream: %v", err)
			}
			log.Printf("Receieved: %v\n", res)
		}
	}()

	<-waitc
}