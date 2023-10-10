package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	// server内の関数に接続
	// 接続不備があるときにcontextを使用する
	// 生成した二つのchannelをserver側に渡している(streamを通している)
	// streamを通したchannel操作の関数が返ってくる
	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	//受信と送信を同時に行うためgoroutineを採用→channelでmain()が終わるのを防いでいる
	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}

		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}

		// send方向のstreamを閉じる
		// からのrequestを送ることでerrorを起こしてstreamを閉じる
		// 両方向のstream操作を複数回行っているため、client streamのようにresponseが返ると同時にclient側のsendStreamを閉じることができない
		// このstreamを通じてchannel通信が行われる
		stream.CloseSend()
	}()

	go func() {
		for {
			// serverからgoroutineで受信している
			res, err := stream.Recv()

			// serverからのstreamが途絶えてしまったらcontextがerrorをはいてRecv()を中断さえる
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Problem while reading server stream: %v\n", err)
			}

			log.Printf("Received a new maximum: %d\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}