package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	// server内の関数に接続
	// 接続不備があるときにcontextを使用する
	// 生成した二つのchannelをserver側に渡している(streamを通している)
	// streamを通したchannel操作の関数が返ってくる
	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		// gorouine(関数内で送信)
		// goroutine越しにrequestを返している
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	// client側のsendStreamを閉じて、最後のresponseを受け取る
	// unaryのようにC.sum(request)で送信受信を同時に行っていないので受信数が一つなのにgoroutineを用いてresponseを受信している
	// responseの受信もgoroutine通信を利用している
	// goroutine通信(内部で受信している)
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}