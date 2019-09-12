package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
)

func main() {

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer con.Close()

	client := calculatorpb.NewCalculatorServiceClient(con)

	// doUnary(client)
	// doServerStreaming(client)
	doClientStreaming(client)
}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{Num1: 3, Num2: 10}

	sum, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("failed on server: %v", err)
	}

	log.Printf("Response from server: %v", sum)
}

func doServerStreaming(client calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculatePrimeNumberDecompositionRequest{
		Number: 120,
	}

	resultStream, err := client.CalculatePrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call CalculatePrimeNumberDecomposition: %v", err)
	}

	for {
		val, err := resultStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive from stream: %v",err)
		}

		log.Printf("Response: %v", val)
	}
}

func doClientStreaming(client calculatorpb.CalculatorServiceClient) {
	log.Println("Sending a call to ComputeAverage...")

	values := []int32{1, 2, 3, 4}

	stream, err := client.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("failed while calling ComputeAverage: %v", err)
	}

	for _, val := range values {
		req := calculatorpb.ComputeAverageRequest{Number:val}
		err := stream.Send(&req)
		if err != nil {
			log.Fatalf("failed while sending to the stream: %v", err)
		}
		log.Printf("Message sent to the stream: %v\n", val)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed while closing stream: %v", err)
	}

	log.Printf("Response: %v\n",resp)

}
