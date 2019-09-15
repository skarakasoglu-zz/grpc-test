package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"net"
)

type server struct{}

func (s *server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	log.Println("Client invoked to FindMaximum")
	max, flag := int32(0), false

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("failed to receive: %v", err)
			return err
		}

		max = func()int32 {
			if req.GetNumber() > max || !flag {
				flag = true
				return req.GetNumber()
			} else {
				return max
			}
		}()
		res := &calculatorpb.FindMaximumResponse{Maximum:max}
		stream.Send(res)
	}
}

func (s *server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	log.Println("Client invoked to ComputeAverage")
	sum, count := int32(0), int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float32(sum) / float32(count)
			resp := calculatorpb.ComputeAverageResponse{Average: average}

			return stream.SendAndClose(&resp)
		}
		if err != nil {
			log.Fatalf("failed while receiving from client: %v", err)
		}

		sum += req.GetNumber()
		count++
	}
}

func (s *server) CalculatePrimeNumberDecomposition(req *calculatorpb.CalculatePrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_CalculatePrimeNumberDecompositionServer) error {
	log.Println("Client invoked to CalculatePrimeNumberDecomposition")
	num := req.GetNumber()

	k := int32(2)

	for num > 1 {
		if num%k == 0 {
			stream.Send(&calculatorpb.CalculatePrimeNumberDecompositionResponse{PrimeNumber: k})
			num /= k
		} else {
			k++
		}
	}

	return nil
}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Client invoked to Sum function...\n")
	firstNumber := req.GetNum1()
	secondNumber := req.GetNum2()

	sum := firstNumber + secondNumber
	resp := &calculatorpb.SumResponse{Result: sum}

	return resp, nil
}

func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	defer lis.Close()

	serv := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(serv, &server{})

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
