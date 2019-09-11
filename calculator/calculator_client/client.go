package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"log"
)

func main() {

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer con.Close()

	client := calculatorpb.NewCalculatorServiceClient(con)

	req := &calculatorpb.SumRequest{Num1: 3, Num2: 10}

	sum, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("failed on server: %v", err)
	}

	log.Printf("Response from server: %v", sum)
}
