package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	// 関数をURL化して、requestをprotocol Bufferに変換した後にserverに送る
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 1,
		SecondNumber: 1,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}