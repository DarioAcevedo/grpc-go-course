package main

import (
	"context"
	"grpc-go-course/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cx, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cx.Close()
	c := calculatorpb.NewCalulatorServiceClient(cx)
	doSumUnary(c)
}


func doSumUnary(c calculatorpb.CalulatorServiceClient) {
	req := calculatorpb.SumRequest{
		A: int32(5),
		B: int32(8),
	}
	res, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}