package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	// server内の関数に接続
	// contextは接続不備があって中断するときに使用する
	// 全部の通信方法で統一することで書きやすくする
	// serverにchannelとrequestを渡している(streamを通している)
	// streamを通したchannel操作の関数が返ってくる
	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Printf("error while calling Primes: %v\n", err)
	}

	for {
		// goroutine(関数内で受信)
		// goroutine(stream)越しにresponseを受け取っている
		res, err := stream.Recv()

		// serverからのstreamが途絶えてしまったらcontextがerrorをはいてRecv()を中断さえる
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", res.Result)
	}
}

//client側は返ってくるstreamに帰属する→sreamを駆使してsendstream, recvstreamを操作する