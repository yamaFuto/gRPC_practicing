package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	// serverに二つのchannelを送信している
	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			// serverにgoroutineで送信している
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		// send方向のstreamを閉じる
		// からのrequestを送ることでerrorを起こしてstreamを閉じる
		// 同時に複数のgoroutineを実行することで両方向のstreamを同時に操作するため、streamを関数の終了と同時に終了する前に強制的に終了させることでclient側のsendStreamを終わらせている
		// このstreamを通じてchannel通信が行われる
		stream.CloseSend()
	}()

	go func() {
		for {
			// serverからgoroutineで受信している
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}